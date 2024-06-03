package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
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

	supply.SourceName = c.PostForm("sourcename")
	supply.Type = CheckType(c.PostForm("type"))
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		log.Warning.Println(err)
	}
	supply.Price = price
	supply.SourceDetail = c.PostForm("sourcedetail")

	imgFile, _ := c.FormFile("imgfile")
	finalFileName := ""
	if imgFile != nil {
		// 创建随机文件名
		fileNameUid := uuid.New()
		// fileNames := strings.Split(testamentJustifyFile.Filename, " ")
		log.Info.Println(imgFile.Filename)
		fileNames := strings.Split(imgFile.Filename, ".")
		finalFileName = fmt.Sprintf("%v.%v", fileNameUid.String(), fileNames[len(fileNames)-1])

		err := c.SaveUploadedFile(imgFile, "../file/imgfile/"+finalFileName)

		if err != nil {
			log.Warning.Println("保存供应商图片文件失败", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "保存供应商图片失败",
			})
			return
		}
	}
	supply.ImgUrl = finalFileName

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

func CheckType(types string) int {
	switch types {
	case "殡仪馆":
		return 1
	case "墓地":
		return 2
	case "火化方式":
		return 3
	case "骨灰盒":
		return 4
	}
	log.Warning.Println(types, "没有找到对应的类别，请维护!")
	return -1
}
