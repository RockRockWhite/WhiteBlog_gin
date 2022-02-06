package middlewares

import (
	"gin/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 计算处理请求消耗时间
		latency := time.Since(startTime)

		// 记录日志
		logger.Logger().Infof("| [request] | %3d | %9v | %15s | %s | %s |", c.Writer.Status(), latency, c.ClientIP(), c.Request.Method, c.Request.RequestURI)
	}
}
