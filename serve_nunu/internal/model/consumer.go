package model

import "gorm.io/gorm"

type Consumer struct {
	gorm.Model
	User_type int    `json:"user_type"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
}

func (m *Consumer) TableName() string {
	return "consumer"
}
