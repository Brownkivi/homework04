package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserId  string
	PostId  string
}

func (Comment) TableName() string {
	return "comments"
}
