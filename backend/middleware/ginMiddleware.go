package middleware

import (
	"dyd/log"
	"dyd/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 验证
func JWTAuthMiddleware() func(c *gin.Context) {
	// token放在Authorization中，使用Bearer开头
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "无权限访问！",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, ":", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "token格式错误",
			})
			c.Abort()
			return
		}

		mc, err := utils.DeruptUserToken(parts[1])
		if err != nil {
			log.Warning.Println("token解析错误", parts[1], err)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "token解析错误",
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}

// CorsMiddleware 解决跨域
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
