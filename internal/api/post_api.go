package api

import (
	"BlodWeb/internal/dao"
	"BlodWeb/internal/model"
	"BlodWeb/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetInt("userID")
	insertPost := model.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  userID,
	}

	if err := dao.CreatePost(&insertPost); err != nil {
		response.Error(c, http.StatusBadRequest, "创建文章失败")
		return
	}

	response.Success(c, "创建文章成功")

}

func GetPostById(c *gin.Context) {
	var post struct {
		Id string `json:"id"`
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	id, err := strconv.ParseInt(post.Id, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if res, err := dao.GetPostById(id); err != nil {
		response.Success(c, res)
		return
	}
	response.Error(c, http.StatusBadRequest, "获取失败")
}

func GetAllPost(c *gin.Context) {
	res, err := dao.GetAllPost()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "获取失败")
	}
	response.Success(c, res)
}

func UpdatePost(c *gin.Context) {
	var post struct {
		Id      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	id, err := strconv.ParseInt(post.Id, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	res, err := dao.GetPostById(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "更新失败")
		return
	}
	userID := c.GetInt("userID")
	if userID != res.UserId {
		response.Error(c, http.StatusBadRequest, "权限不足")
		return
	}
	insertPost := model.Post{
		Title:   post.Title,
		Content: post.Content,
	}
	insertPost.ID = uint(id)
	err = dao.UpdatePostById(&insertPost)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "更新失败")
		return
	}
	response.Success(c, "更新成功")

}

func DeletePost(c *gin.Context) {
	var post struct {
		Id string `json:"id"`
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	id, err := strconv.ParseInt(post.Id, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	res, err := dao.GetPostById(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "更新失败")
		return
	}
	userID := c.GetInt("userID")
	if userID != res.UserId {
		response.Error(c, http.StatusBadRequest, "权限不足")
		return
	}
	if err := dao.DeletePostById(id); err != nil {
		response.Error(c, http.StatusBadRequest, "删除失败")
		return
	}
	response.Success(c, "删除成功")
}
