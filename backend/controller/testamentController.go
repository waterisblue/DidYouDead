package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/middleware"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestamentControllerRegister() {
	engine := getEngine()

	loginAfter := engine.Group("/loginafter").Use(middleware.JWTAuthMiddleware())
	loginAfter.POST("/uploadtestment", uploadTestament)
	loginAfter.GET("/gettestament", getTestament)
	loginAfter.POST("/updatetestamentactive", updateTestamentActive)
}

// 上传遗嘱
func uploadTestament(c *gin.Context) {
	username, _ := c.Get("username")

	testamentDetail, _ := c.GetPostForm("testamentDetail")
	testamentStyle, _ := c.GetPostForm("testamentStyle")
	testamentName, _ := c.GetPostForm("testamentName")

	log.Info.Println(username, "打算建立一份", testamentStyle)
	testamentJustifyFile, _ := c.FormFile("testamentJustifyFile")
	finalFileName := ""
	if testamentJustifyFile != nil {
		// 创建随机文件名
		fileNameUid := uuid.New()
		// fileNames := strings.Split(testamentJustifyFile.Filename, " ")
		log.Info.Println(testamentJustifyFile.Filename)
		fileNames := strings.Split(testamentJustifyFile.Filename, ".")
		finalFileName = fmt.Sprintf("%v.%v", fileNameUid.String(), fileNames[len(fileNames)-1])

		err := c.SaveUploadedFile(testamentJustifyFile, "../file/testament/"+finalFileName)

		if err != nil {
			log.Warning.Println(username, "保存遗嘱文件失败", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "保存遗嘱文件失败",
			})
			return
		}

	}

	go dao.SaveTestament(username.(string), testamentDetail, testamentStyle, finalFileName, testamentName)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传遗嘱成功",
	})

}

func getTestament(c *gin.Context) {
	username, _ := c.Get("username")

	testamentSlice := dao.GetTestamentByUserName(username.(string))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "遗嘱查询成功",
		"data": gin.H{
			"testaments": testamentSlice,
		},
	})
}

func updateTestamentActive(c *gin.Context) {
	username, _ := c.Get("username")
	var testaments []entity.Testament
	err := c.ShouldBind(&testaments)
	if err != nil {
		log.Warning.Println("参数校验失败", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 304,
			"msg":  "参数校验失败",
		})
		return
	}

	for _, v := range testaments {
		log.Info.Println(username, "正在修改遗嘱", v.TestamentName)
		go dao.UpdateTestamentActiveById(v.Id, v.IsActive)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "遗嘱状态修改成功",
	})
}
