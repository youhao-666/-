package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
}

func (m *Student) TableName() string {
    return "student"
}
