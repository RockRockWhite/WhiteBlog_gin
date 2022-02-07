package services

import (
	"fmt"
	"gin/config"
	"testing"
)

func TestNewArticleRepository(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewUserRepository(true)

	fmt.Println(repository)
}
