package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	User_id string
	Post_id string
}
