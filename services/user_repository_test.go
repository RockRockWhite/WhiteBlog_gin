package services

import (
	"fmt"
	"gin/config"
	"gin/entities"
	"testing"
)

func TestNewArticleRepository(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")

	repository := NewUserRepository(true)

	fmt.Println(repository)
}

func TestUserRepository_AddUser(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")
	repository := NewUserRepository(true)

	repository.AddUser(&entities.User{
		NickName: "Rock3",
		Password: "jksdfjk;sdfjkl;dfjkl;dfkl;",
		Salt:     "dffsdfasdasdffs",
	})
}

func TestUserRepository_GetUser(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")
	repository := NewUserRepository(true)

	user := repository.GetUser(1)
	t.Logf("%+v", user)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")
	repository := NewUserRepository(true)

	// 修改用户信息
	user := repository.GetUser(1)
	user.NickName = "被我修改过了"
	repository.UpdateUser(user)

	user = repository.GetUser(1)
	t.Logf("%+v", user)
}

func TestUserRepository_DeleteUser(t *testing.T) {
	// 初始化配置文件
	config.Init("../config/config.yml")
	repository := NewUserRepository(true)

	// 修改用户信息
	repository.DeleteUser(1)

	if repository.UserExists(1) {
		t.Failed()
	}
}
