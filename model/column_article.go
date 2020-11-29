package model

// ColumnArticle 专栏下的文章（一对多）
type ColumnArticle struct {
	ID        uint `gorm:"primarykey"`
	ColumnID  uint
	ArticleID uint
}
