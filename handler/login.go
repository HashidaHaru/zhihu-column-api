package handler

import (
	"zhihu-column-api/dto"
	"zhihu-column-api/service"
	"zhihu-column-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 登陆
func Login(c *gin.Context) {
	var v dto.LoginDto
	if err := c.ShouldBindJSON(&v); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	u, err := service.FindUserByUsername(v.Username)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(v.Password))
	if err != nil {
		errorR(c, bcryptErrCode, err)
		return
	}
	// 校验通过
	token, err := utils.SignToken(u.ID)
	if err != nil {
		errorR(c, tokenErrCode, err)
		return
	}
	successR(c, token)
}
