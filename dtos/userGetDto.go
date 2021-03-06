package dtos

import (
	"gin/entities"
)

// UserGetDto 获得用户Dto
type UserGetDto struct {
	Username    string // 昵称
	Email       string // 邮箱
	VerifyState bool   // 邮箱验证状态
	Telephone   string // 手机号码
	IsAdmin     bool   // 是否管理员
	AvatarUrl   string // 头像链接
}

// ParseUserEntity 将entity转换为GetDto
func ParseUserEntity(user *entities.User) *UserGetDto {
	return &UserGetDto{
		Username:    user.Username,
		Email:       user.Email,
		VerifyState: user.VerifyState,
		Telephone:   user.Telephone,
		IsAdmin:     user.IsAdmin,
		AvatarUrl:   user.AvatarUrl,
	}
}
