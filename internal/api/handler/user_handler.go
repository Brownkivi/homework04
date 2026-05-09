package handler

import (
	"BlodWeb/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录接口
func Login(c *gin.Context) {
	// 模拟账号密码校验，实际走数据库
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	// 模拟查询数据库用户ID
	userID := uint64(10001)

	// 签发JWT
	token, err := jwt.GenerateToken(userID, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "生成Token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

// 需要登录才能访问的接口
func GetUserInfo(c *gin.Context) {
	// 从上下文获取登录用户信息
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"userID":   userID,
		"username": username,
	})
}
