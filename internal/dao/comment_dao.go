package dao

import (
	"BlodWeb/configs"
	"BlodWeb/internal/model"
)

func CreateComment(comment *model.Comment) error {
	var err error
	err = configs.DB.Create(comment).Error
	return err
}

func GetAllComment(id int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := configs.DB.Where("post_id = ?", id).Find(&comments).Error
	return comments, err
}
