package main

import (
	"BlodWeb/configs"
	"BlodWeb/internal/config"
	"BlodWeb/internal/router"

	"go.uber.org/zap"
)

func main() {
	//数据库
	config.InitDB()
	//日志
	config.InitZap()
	//jwt
	r := router.InitRouter()
	//启动服务
	configs.Logger.Info("服务启动")
	err := r.Run(":8080")
	if err != nil {
		configs.Logger.Error("启动失败", zap.Error(err))
		panic("服务启动失败")
	}
}
