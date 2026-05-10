package middleware

import (
	"BlodWeb/configs"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 记录所有接口请求日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 执行接口
		c.Next()

		// 接口结束后打印日志
		cost := time.Since(start)
		status := c.Writer.Status()

		configs.Logger.Info("GIN API",
			zap.String("path", path),
			zap.String("method", method),
			zap.Int("status", status),
			zap.Duration("cost", cost),
			zap.String("ip", c.ClientIP()),
		)
	}
}
