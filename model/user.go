package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username string  `gorm:"type:VARCHAR(20)" json:"username"`
	Password string  `gorm:"type:VARCHAR(300)" json:"password"`
	Email    *string `gorm:"type:VARCHAR(100);unique;" json:"email"`
	Nickname *string `gorm:"type:VARCHAR(100)" json:"nickname"`
	Avatar   *string `gorm:"type:VARCHAR(300)" json:"avatar"`
}
