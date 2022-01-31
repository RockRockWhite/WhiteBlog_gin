package main

import (
	"gin/config"
	"gin/routers"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	config.Init("./config/config.yml")

	// 初始化并运行路由
	router := routers.InitApiRouter()
	_ = router.Run(viper.GetString("HttpServer.Port"))
}
