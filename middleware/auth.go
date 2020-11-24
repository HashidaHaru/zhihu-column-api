package middleware

import (
	"net/http"
	"zhihu-column-api/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuth 令牌验证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Authorization 不能为空",
			})
			c.Abort()
			return
		}
		_, err := utils.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()

		// after request

	}
}
