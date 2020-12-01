package service

import (
	"zhihu-column-api/db"
	"zhihu-column-api/model"
)

// ColumnCreate 创建专栏
func ColumnCreate(column *model.Column) error {
	return db.DB.Create(column).Error
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

// ColumnDiscover 发现专栏
func ColumnDiscover(page int) ([]model.Column, error) {
	list := make([]model.Column, 0)
	err := db.DB.Find(&list).Limit(10).Offset(page).Error
	return list, err
}
