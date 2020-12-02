package dto

// ArticleCreate 创建文章
type ArticleCreate struct {
	HeaderImage string `json:"header_image" binding:"required"`
	Short       string `json:"short" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ColumnID    uint   `json:"column_id" binding:"required"`
}

// ArticleDetail 文章详情
type ArticleDetail struct {
	ArticleID uint `json:"article_id" binding:"required"`
}

// ArticleList 文章列表
type ArticleList struct {
	Page int `json:"page" binding:"required"`
}

// ArticleModify 修改文章
type ArticleModify struct {
	ArticleID   uint   `json:"article_id" binding:"required"`
	HeaderImage string `json:"header_image" binding:"required"`
	Short       string `json:"short" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Title       string `json:"title" binding:"required"`
}

// ArticleDelete 删除文章
type ArticleDelete struct {
	ArticleID uint `json:"article_id" binding:"required"`
}
