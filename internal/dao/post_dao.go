package dao

import (
	"BlodWeb/configs"
	"BlodWeb/internal/model"
)

func CreatePost(post *model.Post) error {
	var err error
	err = configs.DB.Create(post).Error
	return err
}

func GetPostById(postId int64) (*model.Post, error) {
	var post model.Post
	err := configs.DB.First(&post, postId).Error
	return &post, err
}

func GetAllPost() ([]model.Post, error) {
	var posts []model.Post
	err := configs.DB.Find(&posts).Error
	return posts, err
}

func UpdatePostById(post *model.Post) error {
	err := configs.DB.Model(post).Where("id = ?", post.ID).Select("title", "content").Updates(post).Error
	return err
}

func DeletePostById(postId int64) error {
	err := configs.DB.Where("id = ?", postId).Delete(&model.Post{}).Error
	return err
}
