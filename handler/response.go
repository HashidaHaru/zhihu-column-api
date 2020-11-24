package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResp 通用回复
type SuccessResp struct {
	Data interface{} `json:"data,omitempty"`
	Code int         `json:"code"`
}

// ErrorResp 错误回复
type ErrorResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func successR(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResp{
		Data: data,
		Code: 0,
	})
}

func errorR(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, ErrorResp{
		Code: code,
		Msg:  err.Error(),
	})
}

// bindJSONErrCode JSON 绑定错误
const bindJSONErrCode = 999

// serviceErrCode 数据库操作错误
const serviceErrCode = 998

// bcrybtErrCode 加密错误
const bcryptErrCode = 997

// tokenErrCode token 签发错误
const tokenErrCode = 996
