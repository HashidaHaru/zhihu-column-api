package model

import "gorm.io/gorm"

// Column 专栏
type Column struct {
	gorm.Model
	Cover       string `gorm:"type:VARCHAR(300)"`
	Description string `gorm:"type:VARCHAR(300)"`
	Title       string `gorm:"type:VARCHAR(100)"`
	AuthorID    uint
}
