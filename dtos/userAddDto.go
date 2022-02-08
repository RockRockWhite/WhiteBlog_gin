package dtos

import (
	"gin/entities"
	"gin/utils"
)

// UserAddDto 创建用户Dto
type UserAddDto struct {
	NickName  string // 昵称
	Password  string // 密码
	Email     string // 邮箱
	Telephone string // 手机号码
	AvatarUrl string // 头像链接
}

// ToEntity 转换成Entity
func (dto *UserAddDto) ToEntity() *entities.User {
	salt := utils.GenerateSalt()
	passwordHash := utils.EncryptPasswordHash(dto.Password, salt)

	return &entities.User{
		NickName:     dto.NickName,
		PasswordHash: passwordHash,
		Salt:         salt,
		Email:        dto.Email,
		VerifyState:  false,
		Telephone:    dto.Telephone,
		IsAdmin:      false,
	}
}
