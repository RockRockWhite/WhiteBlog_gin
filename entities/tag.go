package entities

import "gorm.io/gorm"

// Tag 博文标签实体类
type Tag struct {
	gorm.Model
	ArticleId uint   // 博文Id
	Name      string // 标签名称
}

func (Tag) TableName() string {
	return "a_tags"
}
