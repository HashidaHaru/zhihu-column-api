package model

import "gorm.io/gorm"

// Article 文章
type Article struct {
	gorm.Model
	Short   string `gorm:"type:VARCHAR(300)"`
	Content string `gorm:"type:VARCHAR(2000)"`
	Title   string `gorm:"type:VARCHAR(100)"`
}
