package configs

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// JWT配置
var JwtConfig = struct {
	SecretKey     string        // 密钥
	ExpireTime    time.Duration // 过期时长
	RefreshExpire time.Duration // 刷新令牌过期
}{
	SecretKey:     "your-enterprise-secret-key-123456",
	ExpireTime:    2 * time.Hour,      // 访问令牌2小时过期
	RefreshExpire: 7 * 24 * time.Hour, // 刷新令牌7天
}

// 加密配置
var Key = []byte("12345678901234567890123456789012") //AES密钥

// 全局DB
var DB *gorm.DB

// 白名单：不需要登录拦截的接口
var WhiteList = []string{
	"/api/user/register",
	"/api/user/login",
}

// 全局日志，项目任何地方直接用
var Logger *zap.Logger
