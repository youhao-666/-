package model

import "gorm.io/gorm"

type People struct {
	gorm.Model
}

func (m *People) TableName() string {
    return "people"
}
