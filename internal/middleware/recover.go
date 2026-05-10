package middleware

import (
	"BlodWeb/internal/common/exception"
	"BlodWeb/utils/response"

	"github.com/gin-gonic/gin"
)

// RecoverMiddleware 全局统一错误处理
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 1. 如果是 业务错误
				if bizErr, ok := err.(*exception.BizError); ok {
					response.Error(c, bizErr.Code, bizErr.Msg)
					c.Abort()
					return
				}

				// 2. 系统错误（未知错误）
				response.Error(c, 500, "服务器异常")
				c.Abort()
			}
		}()

		c.Next()
	}
}
