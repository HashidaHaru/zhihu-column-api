package service

import (
	"errors"
	"zhihu-column-api/db"
	"zhihu-column-api/model"
)

// ArticleCreate 创建文章
func ArticleCreate(m *model.Article) error {
	return db.DB.Create(m).Error
}

// Detail 文章详情
type Detail struct {
	HeaderImage    string `gorm:"type:VARCHAR(300)" json:"header_image"`
	Short          string `gorm:"type:VARCHAR(300)" json:"short"`
	Content        string `gorm:"type:VARCHAR(2000)" json:"content"`
	Title          string `gorm:"type:VARCHAR(100)" json:"title"`
	AuthorID       uint   `json:"author_id"`
	AuthorNickname string `json:"author_nickname"`
	AuthorAvatar   string `json:"author_avatar"`
	ColumnID       uint   `json:"column_id"`
}

// ArticleList 文章列表
func ArticleList(id uint, page int) ([]model.Article, error) {
	list := make([]model.Article, 0)
	err := db.DB.Where("author_id = ?", id).Limit(10).Offset((page - 1) * 10).Find(&list).Error
	return list, err
}

// ArticleDetail 文章详情
func ArticleDetail(id uint) (Detail, error) {
	var a Detail
	sql := db.DB.Table("articles").
		Select("articles.*,users.avatar AS author_avatar,users.nickname AS author_nickname").
		Joins("LEFT JOIN users ON articles.author_id = users.id ").
		Where("articles.id = ? AND articles.deleted_at IS NULL", id).
		Find(&a)
	affected := sql.RowsAffected
	if affected == 0 {
		return a, errors.New("article not found")
	}
	err := sql.Error
	return a, err
}

// ArticleModify 修改文章
func ArticleModify(id uint, m model.Article) error {
	return db.DB.Model(&model.Article{}).Where("id = ?", id).Updates(m).Error
}

// ArticleDelete 删除文章
func ArticleDelete(id uint) error {
	return db.DB.Delete(&model.Article{}, id).Error
}
