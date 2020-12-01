package handler

import (
	"zhihu-column-api/dto"
	"zhihu-column-api/model"
	"zhihu-column-api/service"
	"zhihu-column-api/utils"

	"github.com/gin-gonic/gin"
)

// ArticleCreate 创建文章
func ArticleCreate(c *gin.Context) {
	var d dto.ArticleCreate
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	claims, _ := c.Get("claims")
	m := claims.(*utils.MyCustomClaims)
	a := model.Article{
		HeaderImage: d.HeaderImage,
		Short:       d.Short,
		Content:     d.Content,
		Title:       d.Title,
		ColumnID:    d.ColumnID,
		AuthorID:    m.UserID,
	}

	if err := service.ArticleCreate(&a); err != nil {
		errorR(c, serviceErrCode, err)
		return
	}

	successR(c, a.ID)
}

// ArticleDetail 文章详情
func ArticleDetail(c *gin.Context) {
	var d dto.ArticleDetail
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	a, err := service.ArticleDetail(d.ArticleID)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, a)
}

// ArticleModify 修改文章
func ArticleModify(c *gin.Context) {

}

// ArticleDelete 删除文章
func ArticleDelete(c *gin.Context) {

}
