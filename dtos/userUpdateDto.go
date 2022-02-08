package dtos

import (
	"gin/entities"
	"gin/utils"
)

// UserUpdateDto 添加用户Dto
type UserUpdateDto struct {
	NickName  string // 昵称
	Password  string // 密码
	Email     string // 邮箱
	Telephone string // 手机号码
	IsAdmin   bool   // 是否管理员
	AvatarUrl string // 头像链接
}

// UserUpdateDtoFromEntity 从entity转换UpdateDto
func UserUpdateDtoFromEntity(user *entities.User) *UserUpdateDto {
	return &UserUpdateDto{
		NickName:  user.NickName,
		Password:  "",
		Email:     user.Email,
		Telephone: user.Telephone,
		IsAdmin:   user.IsAdmin,
		AvatarUrl: user.AvatarUrl,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *UserUpdateDto) ApplyUpdateToEntity(entity *entities.User) {
	entity.NickName = dto.NickName
	// 计算密码盐值
	if dto.Password != "" {
		entity.PasswordHash = utils.EncryptPasswordHash(dto.Password, entity.Salt)
	}
	entity.Email = dto.Email
	entity.Telephone = dto.Telephone
	entity.IsAdmin = dto.IsAdmin
	entity.AvatarUrl = dto.AvatarUrl
}
