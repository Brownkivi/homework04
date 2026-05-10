package api

import (
	"BlodWeb/internal/dao"
	"BlodWeb/internal/model"
	"BlodWeb/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var param struct {
		PostId  string `json:"postId"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&param); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetString("userID")
	insertComment := model.Comment{
		Content: param.Content,
		UserId:  userID,
		PostId:  param.PostId,
	}
	if err := dao.CreateComment(&insertComment); err != nil {
		response.Error(c, http.StatusBadRequest, "发布评论失败")
		return
	}

	response.Success(c, "发布评论成功")
}

func GetAllComment(c *gin.Context) {
	var param struct {
		PostId string `json:"postId"`
	}
	if err := c.ShouldBindJSON(&param); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	id, err := strconv.ParseInt(param.PostId, 10, 64)
	res, err := dao.GetAllComment(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "获取评论失败")
	}
	response.Success(c, res)

}
