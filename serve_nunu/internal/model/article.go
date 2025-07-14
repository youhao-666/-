package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserId  uint   `json:"user_id" gorm:"column:user_id"`
	Content string `json:"content" gorm:"column:content"`
	Title   string `json:"title" gorm:"column:title"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"column:article_id"`
	TagID     uint `gorm:"column:tag_id"`
}

func (m *Article) TableName() string {
	return "article"
}
