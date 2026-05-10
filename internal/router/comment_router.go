package router

import (
	"BlodWeb/internal/api"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/comment")
	{
		userGroup.POST("/createComment", api.CreateComment)
		userGroup.POST("/getAllComment", api.GetAllComment)
	}
}
