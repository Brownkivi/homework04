package dao

import (
	"BlodWeb/configs"
	"BlodWeb/internal/model"
)

func SelectByUsername(user *model.User, username string) error {
	var err error
	err = configs.DB.Where("name = ?", username).First(user).Error
	return err
}

func SelectCountByUsername(user *model.User, username string) int64 {
	var count int64
	configs.DB.Where("name = ?", username).First(user).Count(&count)
	return count
}
