package router

import (
	"BlodWeb/internal/api"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", api.Register)
		userGroup.POST("/login", api.Login)
		userGroup.POST("/modifyPassword", api.ModifyPassword)
		userGroup.POST("/getUserInfo", api.GetUserInfo)
	}
}
