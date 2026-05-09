package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserId  int
}

func (Post) TableName() string {
	return "posts"
}
