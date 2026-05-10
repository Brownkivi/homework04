package middleware

import (
	"BlodWeb/configs"
	"BlodWeb/utils/jwt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		//白名单放行
		path := c.Request.URL.Path
		if isWhiteList(path) {
			c.Next()
			return
		}

		// 1. 从Header获取Token: Authorization: Bearer xxx
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登录或Token格式错误",
			})
			c.Abort()
			return
		}

		tokenStr := authHeader[7:]
		// 2. 解析Token
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token已过期或无效",
			})
			c.Abort()
			return
		}

		// 3. 将用户信息存入Gin上下文，后续接口可直接获取
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// 判断当前路径是否在白名单
func isWhiteList(path string) bool {
	for _, p := range configs.WhiteList {
		if path == p {
			return true
		}
	}
	return false
}
