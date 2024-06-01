package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SupplyControllerRegister() {
	engine := getEngine()

	loginAfter := engine.Group("/loginafter").Use(middleware.JWTAuthMiddleware())
	loginAfter.POST("/getsupply", GetSupply)
	loginAfter.POST("/addsupply", AddSupply)
	loginAfter.POST("/deletesupply", DeleteSupply)
}

func GetSupply(c *gin.Context) {
	supplyList := dao.GetAllSupply()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功！",
		"data": supplyList,
	})
}

func AddSupply(c *gin.Context) {
	var supply entity.SupplyEntity
	err := c.ShouldBind(&supply)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数校验失败！",
		})
		log.Warning.Println(err)
	}
	go dao.AddSupply(supply)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "增加成功！",
	})
}

func DeleteSupply(c *gin.Context) {
	var supply entity.SupplyEntity
	err := c.ShouldBind(&supply)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数校验失败！",
		})
		log.Warning.Println(err)
	}
	go dao.DeleteSupply(supply)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功！",
	})
}
