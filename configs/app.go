package configs

import "time"

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
