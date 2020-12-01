package dto

// ColumnCreate 创建专栏
type ColumnCreate struct {
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// ColumnDetail 专栏详情
type ColumnDetail struct {
	ColumnID uint `json:"column_id"`
}

// ColumnModify 修改专栏
type ColumnModify struct {
	ColumnID    uint   `json:"column_id"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// ColumnDiscover 发现专栏
type ColumnDiscover struct {
	Page int `json:"page"`
}
