package config

import (
	"BlodWeb/configs"
	"BlodWeb/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	dbPath := "D:/Softstore/Sqlite/sqliteData/Blog.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	// 🔥 修复：必须打开建表，否则没有表无法插入数据
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		panic("建表失败：" + err.Error())
	}
	println("✅ 博客数据库表创建成功！")
	configs.DB = db
}
