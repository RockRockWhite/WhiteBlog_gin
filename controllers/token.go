package controllers

import (
	"fmt"
	"gin/dtos"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// CreateToken 创建token
func CreateToken(c *gin.Context) {
	var tokenDto dtos.TokenAddDto

	if err := c.ShouldBind(&tokenDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 查找用户
	if !userRepository.UsernameExists(tokenDto.Username) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "invalid username or password",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user := userRepository.GetUserByNickName(tokenDto.Username)

	// 验证密码
	if !utils.ValifyPasswordHash(tokenDto.Password, user.Salt, user.PasswordHash) {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "invalid username or password",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	token, err := utils.GenerateJwtToken(&utils.JwtClaims{
		Id:          user.ID,
		NickName:    user.Username,
		Email:       user.Email,
		VerifyState: user.VerifyState,
		Telephone:   user.Telephone,
		IsAdmin:     user.IsAdmin,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to generate JwtToken"))
	}

	c.JSON(http.StatusCreated, struct {
		Token string
	}{Token: token})
}
