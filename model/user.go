package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR(20)"`
	Password string `gorm:"type:VARCHAR(300)"`
	Email    string `gorm:"type:VARCHAR(100);unique;not null"`
	Nickname string `gorm:"type:VARCHAR(100)"`
	Avatar   string `gorm:"type:VARCHAR(300)"`
}
