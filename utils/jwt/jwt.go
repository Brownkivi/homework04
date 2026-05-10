package jwt

import (
	"BlodWeb/configs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义JWT载荷
type CustomClaims struct {
	UserID               uint64 `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准声明：过期、签发时间等
}

// 生成Token
func GenerateToken(userID uint64, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.JwtConfig.ExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 使用HS256签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JwtConfig.SecretKey))
}

// 解析Token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(configs.JwtConfig.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token无效")
}
