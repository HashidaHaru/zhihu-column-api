package handler

import (
	"zhihu-column-api/dto"
	"zhihu-column-api/model"
	"zhihu-column-api/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
	}

	if err := service.UserAdd(&user); err != nil {
		errorR(c, serviceErrCode, err)
		return
	}

	successR(c, user.ID)
}
