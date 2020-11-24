package router

import (
	"zhihu-column-api/handler"
	"zhihu-column-api/middleware"

	"github.com/gin-gonic/gin"
)

// RunHTTP http 服务
func RunHTTP() {
	r := gin.Default()
	r.Use(middleware.Cros())

	// 登录接口
	r.POST("/user/login", handler.Login)
	// 注册接口
	r.POST("/user/signup", handler.Signup)
	r.Run(":8081")
}
