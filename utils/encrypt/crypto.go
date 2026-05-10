package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// AES-GCM 加密
func Encrypt(plainText string, key []byte) (string, error) {
	// 1. 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 2. 获取 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 3. 生成随机 nonce（必须每次加密都不一样）
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 4. 加密（自动追加认证标签）
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	// 5. 返回 Base64 字符串（方便传输/存储）
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// AES-GCM 解密
func Decrypt(cipherText string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherData := data[:nonceSize], data[nonceSize:]

	// 解密 + 验证数据完整性
	plainText, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
