package dto

// ColumnCreate 创建专栏
type ColumnCreate struct {
	Cover       string `json:"cover" binding:"required"`
	Description string `json:"description" binding:"required"`
	Title       string `json:"title" binding:"required"`
}

// ColumnList 查看我的专栏列表
type ColumnList struct {
	Page int `json:"page" binding:"required"`
}

// ColumnModify 修改专栏
type ColumnModify struct {
	ColumnID    uint   `json:"column_id" binding:"required"`
	Cover       string `json:"cover" binding:"required"`
	Description string `json:"description" binding:"required"`
	Title       string `json:"title" binding:"required"`
}

// ColumnDiscover 发现专栏
type ColumnDiscover struct {
	Page int `json:"page" binding:"required"`
}
