package controllers

import (
	"fmt"
	"gin/dtos"
	"gin/services"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var userRepository *services.UserRepository

// InitUserController 初始化用户Controller
func InitUserController() {
	userRepository = services.NewUserRepository(true)
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var userDto dtos.UserAddDto

	if err := c.ShouldBind(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity := userDto.ToEntity()
	userRepository.AddUser(entity)

	c.JSON(http.StatusCreated, dtos.ParseUserEntity(entity))
}

// GetUser 获得用户
func GetUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	user := userRepository.GetUser(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParseUserEntity(user))
}

// PatchUser 修改用户
func PatchUser(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	user := userRepository.GetUser(uint(id))

	// 获得patchJson
	patchJson, getRawDataErr := c.GetRawData()
	if getRawDataErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 应用patch
	dto := dtos.UserUpdateDtoFromEntity(user)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(user)

	// 更新数据库
	userRepository.UpdateUser(user)

	c.Status(http.StatusNoContent)
}

// PutUser 替换用户
func PutUser(c *gin.Context) {
	// TODO: 待完成
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !userRepository.UserExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("User id %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	userRepository.DeleteUser(uint(id))
	c.Status(http.StatusNoContent)
}
