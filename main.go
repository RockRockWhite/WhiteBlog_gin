package main

import (
	"gin/config"
	"gin/routers"
	"gin/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	config.Init("./config/config.yml")

	// 初始化logger
	utils.InitLogger(viper.GetString("Logger.LogFile"), logrus.DebugLevel, "2006-01-02 15:04:05")
	utils.Logger().Infof("| [service] | ***** Service started ***** |")
	defer utils.Logger().Infof("| [service] | ***** Service stoped ***** |")

	// 初始化并运行路由
	router := routers.InitApiRouter()

	_ = router.Run(viper.GetString("HttpServer.Port"))
}
