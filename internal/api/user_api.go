package api

import (
	"BlodWeb/configs"
	"BlodWeb/internal/dao"
	"BlodWeb/internal/model"

	"BlodWeb/utils/encrypt"
	"BlodWeb/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 登录接口
func Login(c *gin.Context) {
	// 模拟账号密码校验，实际走数据库
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		configs.Logger.Error("参数错误", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	var user model.User
	if err := dao.SelectByUsername(&user, req.Username); err != nil {
		configs.Logger.Error("用户名或密码错误", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	//解密密码
	password, err := encrypt.Decrypt(user.Password, configs.Key)
	if err != nil {
		configs.Logger.Error("加密失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "加密失败"})
		return
	}

	//密码错误
	if password != req.Password {
		configs.Logger.Error("密码错误", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	userID := uint64(user.ID)

	// 签发JWT
	token, err := jwt.GenerateToken(userID, req.Username)
	if err != nil {
		configs.Logger.Error("生成Token失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "生成Token失败"})
		return
	}

	configs.Logger.Info("登录成功", zap.String("userId", strconv.FormatUint(userID, 10)))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

// 注册接口
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}
	var count int64
	count = dao.SelectCountByUsername(&model.User{}, req.Username)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名已存在"})
		return
	}

	var encodeStr string
	encodeStr, err := encrypt.Encrypt(req.Password, configs.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "加密失败"})
		return
	}
	req.Password = encodeStr

	err = configs.DB.Create(&model.User{
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

// 修改密码
func ModifyPassword(c *gin.Context) {
	// 从上下文获取登录用户信息
	userID, _ := c.Get("userID")
	var req struct {
		OldPassword string `json:"oldPassword"`
		Password    string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}
	//当前登录用户旧密码
	var password string
	if err := configs.DB.Model(&model.User{}).Where("id = ?", userID).Pluck("password", &password).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "修改密码失败"})
		return
	}
	var decodeStr string
	decodeStr, err := encrypt.Decrypt(password, configs.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "加密失败"})
		return
	}
	if decodeStr != req.OldPassword {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
		return
	}
	var newEncodeStr string
	newEncodeStr, err = encrypt.Encrypt(req.Password, configs.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "加密失败"})
		return
	}

	err = configs.DB.Model(&model.User{}).Where("id = ?", userID).Update("password", newEncodeStr).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "修改成功"})
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
