package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:name"`
	Password string
	Email    string
}

func (User) TableNameI() string {
	return "users"
}
