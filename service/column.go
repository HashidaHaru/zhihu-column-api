package service

import (
	"zhihu-column-api/db"
	"zhihu-column-api/model"
)

// ColumnCreate 创建专栏
func ColumnCreate(column *model.Column) error {
	return db.DB.Create(column).Error
}

// ColumnAll 查看我的专栏列表
func ColumnAll(id uint) ([]model.Column, error) {
	list := make([]model.Column, 0)
	err := db.DB.Where("author_id= ? ", id).Find(&list).Error
	return list, err
}

// ColumnDetail 专栏详情
func ColumnDetail(id uint) (*model.Column, error) {
	column := model.Column{}
	column.ID = id
	err := db.DB.First(&column).Error
	return &column, err
}

// ColumnModify 修改专栏
func ColumnModify(id uint, m model.Column) error {
	return db.DB.Model(&model.Column{}).Where("id = ?", id).Updates(m).Error
}

// DiscoverColumn 发现专栏列表
type DiscoverColumn struct {
	Cover          string `json:"cover"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	AuthorID       uint   `json:"author_id"`
	AuthorNickname string `json:"author_nickname"`
	AuthorAvatar   string `json:"author_avatar"`
}

// ColumnDiscover 发现专栏
func ColumnDiscover(page int) ([]DiscoverColumn, error) {
	list := make([]DiscoverColumn, 0)
	err := db.DB.Table("columns").
		Select("columns.*,users.avatar AS author_avatar,users.nickname AS author_nickname").
		Joins("LEFT JOIN users ON columns.author_id = users.id ").
		Limit(10).Offset((page - 1) * 10).Find(&list).Error
	return list, err
}
