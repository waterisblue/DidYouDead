package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FireServiceControllerRegister() {
	engine := getEngine()

	loginAfter := engine.Group("/loginafter").Use(middleware.JWTAuthMiddleware())
	loginAfter.POST("/uploadFireService", uploadFireService)
	loginAfter.POST("/getFireServiceByUserName", getFireServiceByUserName)

}

// 上传火化方式
func uploadFireService(c *gin.Context) {
	var fireService entity.FireService
	err := c.ShouldBind(&fireService)
	UserName, _ := c.Get("username")
	if err != nil {
		log.Warning.Println("uploadFireService 参数校验失败", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 304,
			"msg":  "参数校验失败",
		})
		return
	}
	go dao.InsertFireService(fireService, UserName.(string))
	c.JSON(http.StatusOK, gin.H{
		"code": 304,
		"msg":  "火化方案提交成功！",
	})
	return
}

// 获取火化方式
func getFireServiceByUserName(c *gin.Context) {
	username, _ := c.Get("username")
	exists, fireService := dao.GetFireServiceByUserName(username.(string))

	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "没有查询到遗嘱信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"fireService": fireService,
		},
	})
	return
}
