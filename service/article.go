package service

import (
	"zhihu-column-api/db"
	"zhihu-column-api/model"
)

// ArticleCreate 创建文章
func ArticleCreate(m *model.Article) error {
	return db.DB.Create(m).Error
}

// Detail 文章详情
type Detail struct {
}

// ArticleDetail 文章详情
func ArticleDetail(id uint) (Detail, error) {
	var a Detail

	err := db.DB.Error
	return a, err
}

// ArticleModify 修改文章
func ArticleModify() {

}

// ArticleDelete 删除文章
func ArticleDelete() {}
