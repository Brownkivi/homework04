package router

import (
	"BlodWeb/internal/api/handler"
	"BlodWeb/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 无需登录路由
	public := r.Group("/api")
	{
		public.POST("/login", handler.Login)
	}

	// 需要JWT认证路由
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/user/info", handler.GetUserInfo)
	}

	return r
}
