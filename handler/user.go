package handler

import (
	"zhihu-column-api/dto"
	"zhihu-column-api/model"
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

// Signup 注册
func Signup(c *gin.Context) {
	var v dto.SignupDto
	if err := c.ShouldBindJSON(&v); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(v.Password), 11)
	if err != nil {
		errorR(c, bcryptErrCode, err)
		return
	}

	user := model.User{
		Username: v.Username,
		Password: string(hash),
		Email:    v.Email,
	}

	if err := service.UserAdd(&user); err != nil {
		errorR(c, serviceErrCode, err)
		return
	}

	successR(c, user.ID)
}

// UserDetail 用户详情
func UserDetail(c *gin.Context) {
	claims, _ := c.Get("claims")
	cm := claims.(*utils.MyCustomClaims)
	d, err := service.UserDetail(cm.UserID)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return
	}
	successR(c, d)
}

// UserModify 修改用户详情
func UserModify(c *gin.Context) {
	var d dto.UserModify
	if err := c.ShouldBindJSON(&d); err != nil {
		errorR(c, bindJSONErrCode, err)
		return
	}
	claims, _ := c.Get("claims")
	cm := claims.(*utils.MyCustomClaims)
	m := model.User{
		Email:    d.Email,
		Nickname: &d.Nickname,
		Avatar:   &d.Avatar,
	}
	err := service.UserModify(cm.UserID, m)
	if err != nil {
		errorR(c, serviceErrCode, err)
		return

	}
	successR(c, nil)
}
