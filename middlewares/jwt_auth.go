package middlewares

import (
	"gin/dtos"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// JwtAuth JwtToken验证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// 没有传token参数
			c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
				Message:          "Requires key {Authorization} in headers",
				DocumentationUrl: viper.GetString("Document.Url"),
			})

			c.Abort()
			return
		}

		claims, err := utils.ParseJwtToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
				Message:          "Token expired or the other error occurred",
				DocumentationUrl: viper.GetString("Document.Url"),
			})

			c.Abort()
			return
		}

		// Claims写入上下文
		c.Set("claims", claims)
	}
}
