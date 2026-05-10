package router

import (
	"BlodWeb/internal/api"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/post")
	{
		userGroup.POST("/createPost", api.CreatePost)
		userGroup.POST("/getPostById", api.GetPostById)
		userGroup.POST("/getAllPost", api.GetAllPost)
		userGroup.POST("/updatePost", api.UpdatePost)
		userGroup.POST("/deletePost", api.DeletePost)
	}
}
