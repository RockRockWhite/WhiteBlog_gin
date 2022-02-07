package services

import (
	"fmt"
	"gin/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建新用户Repository
func NewUserRepository(autoMigrate bool) *UserRepository {
	Host := viper.GetString("DataBase.Host")
	Port := viper.GetString("DataBase.Port")
	Username := viper.GetString("DataBase.Username")
	Password := viper.GetString("DataBase.Password")
	DBName := viper.GetString("DataBase.DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error open database:%s %s \n", dsn, err))
	}

	// 完成User迁移
	if autoMigrate {
		if err := db.AutoMigrate(&entities.User{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "User", err))
		}
	}

	repository := UserRepository{db}
	return &repository
}
func (repository *UserRepository) AddUser(id uint) *entities.User {
	return nil
}

func (repository *UserRepository) GetUser(id uint) *entities.User {
	return nil
}

func (repository *UserRepository) UpdateUser(id uint) *entities.User {
	return nil
}

func (repository *UserRepository) DeleteUser(id uint) *entities.User {
	return nil
}
