package handler

import (
	"fmt"
	"zhihu-column-api/dto"
	"zhihu-column-api/model"
	"zhihu-column-api/service"
	"zhihu-column-api/utils"

	"github.com/gin-gonic/gin"
)

// ColumnCreate 创建专栏
func ColumnCreate(c *gin.Context) {
	var d dto.ColumnCreate
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	claims, _ := c.Get("claims")
	m := claims.(*utils.MyCustomClaims)
	column := model.Column{
		Cover:       d.Cover,
		Description: d.Description,
		Title:       d.Title,
		AuthorID:    m.UserID,
	}

	if err := service.ColumnCreate(&column); err != nil {
		errorR(c, serviceErrCode, err)
		return
	}

	successR(c, column.ID)

}

// ColumnList 查看我的专栏列表
func ColumnList(c *gin.Context) {
	var d dto.ColumnList
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	claims, _ := c.Get("claims")
	cm := claims.(*utils.MyCustomClaims)
	list, err := service.ColumnList(cm.UserID, d.Page)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, list)
}

// ColumnModify 修改专栏
func ColumnModify(c *gin.Context) {
	var d dto.ColumnModify
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	// todo 检查专栏作者 id 是否为用户 id
	claims, _ := c.Get("claims")
	cm := claims.(*utils.MyCustomClaims)
	column, err := service.ColumnDetail(d.ColumnID)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	if column.AuthorID != cm.UserID {
		errorR(c, authorErrCode, fmt.Errorf("权限不足"))
	}
	m := model.Column{
		Cover:       d.Cover,
		Description: d.Description,
		Title:       d.Title,
	}
	err = service.ColumnModify(d.ColumnID, m)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, nil)
}

// ColumnDiscover 发现专栏
func ColumnDiscover(c *gin.Context) {
	var d dto.ColumnDiscover
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	list, err := service.ColumnDiscover(d.Page)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, list)
}
