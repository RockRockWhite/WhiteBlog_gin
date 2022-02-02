package dtos

import (
	"gin/entities"
)

// StarAddDto 创建博文点赞Dto
type StarAddDto struct{}

// ToEntity 转换成Entity
func (dto *StarAddDto) ToEntity(articleId uint, userId uint) *entities.Star {
	return &entities.Star{
		UserId:    userId,
		ArticleId: articleId,
	}
}
