package dtos

import (
	"gin/entities"
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
		Password:  user.Salt,
		Email:     user.Email,
		Telephone: user.Telephone,
		IsAdmin:   user.IsAdmin,
		AvatarUrl: user.AvatarUrl,
	}
}

// ApplyUpdateToEntity 将Update应用到Entity
func (dto *UserUpdateDto) ApplyUpdateToEntity(entity *entities.User) {
	entity.NickName = dto.NickName
	// TODO: 解决修改密码问题
	entity.PasswordHash = dto.Password
	entity.Email = dto.Email
	entity.Telephone = dto.Telephone
	entity.IsAdmin = dto.IsAdmin
	entity.AvatarUrl = dto.AvatarUrl
}
