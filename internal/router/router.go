package router

import (
	"BlodWeb/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// api
	api := r.Group("/api")
	//全局jwt
	api.Use(middleware.JWTAuth())
	//全局zap日志
	api.Use(middleware.GinLogger())
	//统一异常处理
	api.Use(middleware.RecoverMiddleware())
	UserRoutes(api)

	return r
}
