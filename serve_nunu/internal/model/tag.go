package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string `json:"tagName" gorm:"column:tag_name"`
}

func (m *Tag) TableName() string {
	return "tag"
}
