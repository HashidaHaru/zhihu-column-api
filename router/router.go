package router

import (
	"zhihu-column-api/middleware"

	"github.com/gin-gonic/gin"
)

// RunHTTP http 服务
func RunHTTP() {
	r := gin.Default()
	r.Use(middleware.Cros())

	r.Run(":8080")
}
