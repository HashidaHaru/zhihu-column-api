package model

import "gorm.io/gorm"

// Article 文章
type Article struct {
	gorm.Model
	HeaderImage string `gorm:"type:VARCHAR(300)" json:"header_image"`
	Short       string `gorm:"type:VARCHAR(300)" json:"short"`
	Content     string `gorm:"type:VARCHAR(2000)" json:"content"`
	Title       string `gorm:"type:VARCHAR(100)" json:"title"`
}
