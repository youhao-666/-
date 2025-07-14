package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
