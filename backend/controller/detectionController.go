package controller

import (
	"dyd/dao"
	"dyd/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DetectionControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	loginAfter := engine.Group("/loginafter").Use(middleware.JWTAuthMiddleware())
	loginAfter.POST("/getDetectionData", getDetectionData)
}

func getDetectionData(c *gin.Context) {
	username, _ := c.Get("username")
	_, userdetail := dao.GetUserByAccount(username.(string))
	if !userdetail.Administrator {
		c.JSON(http.StatusForbidden, gin.H{
			"code": http.StatusForbidden,
			"msg":  "您的账号无权限获取此数据",
		})
		return
	}
	detections := dao.GetDetectionByDate()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取数据成功",
		"data": detections,
	})
}
