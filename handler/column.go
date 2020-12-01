package handler

import (
	"zhihu-column-api/dto"
	"zhihu-column-api/model"
	"zhihu-column-api/service"

	"github.com/gin-gonic/gin"
)

// ColumnCreate 创建专栏
func ColumnCreate(c *gin.Context) {
	var d dto.ColumnCreate
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}

	column := model.Column{
		Cover:       d.Cover,
		Description: d.Description,
		Title:       d.Title,
	}

	if err := service.ColumnCreate(&column); err != nil {
		errorR(c, serviceErrCode, err)
		return
	}

	successR(c, column.ID)

}

// ColumnDetail 专栏详情
func ColumnDetail(c *gin.Context) {
	var d dto.ColumnDetail
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	column, err := service.ColumnDetail(d.ColumnID)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, column)
}

// ColumnModify 修改专栏
func ColumnModify(c *gin.Context) {
	var d dto.ColumnModify
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	m := model.Column{
		Cover:       d.Cover,
		Description: d.Description,
		Title:       d.Title,
	}
	err := service.ColumnModify(d.ColumnID, m)
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
