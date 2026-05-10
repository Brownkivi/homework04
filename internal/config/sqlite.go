package config

import (
	"BlodWeb/configs"
	"BlodWeb/internal/model"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	dbPath := "D:/Softstore/Sqlite/Blog.db"

	dbDir := filepath.Dir(dbPath)
	// 判断目录是否存在
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		// 不存在就创建目录（包括多级目录）
		err = os.MkdirAll(dbDir, 0755)
		if err != nil {
			panic("创建数据库目录失败：" + err.Error())
		}
	}

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
