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
	r.Use(middleware.JWTAuth())

	// 创建专栏
	r.POST("/column/create", handler.ColumnCreate)
	// 查看我的专栏信息
	r.POST("/column/detail", handler.ColumnDetail)
	// 修改我的专栏信息
	r.POST("/column/modify", handler.ColumnModify)
	// 发现专栏
	r.POST("/column/discover", handler.ColumnDiscover)

	// // 创建文章
	// r.POST("/article/create", handler.ArticleCreate)
	// // 修改文章
	// r.POST("/article/modify", handler.ArticleModify)
	// // 删除文章
	// r.POST("/article/delete", handler.ArticleDelete)

	// 点赞

	// 评论

	r.Run(":8081")
}
