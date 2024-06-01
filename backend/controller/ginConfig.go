package controller

import (
	log "dyd/log"

	"dyd/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

// InitGin 初始化gin
func init() {
	engine = gin.Default()
}

// 获取gin
func getEngine() (engin *gin.Engine) {
	if engine == nil {
		log.Error.Println("gin engine还未进行初始化，无法获取，请先进行初始化！")
	}
	return engine
}

// StartGin 开启服务器
func StartGin() {
	// 配置跨域中间件
	engine.Use(middleware.CorsMiddleware())
	// 配置静态文件夹
	engine.Use(static.Serve("/static/testament", static.LocalFile("../file/testament", true)))
	UserControllerRegister()
	AuthControllerRegister()
	TestamentControllerRegister()
	FireServiceControllerRegister()
	DetectionControllerRegister()
	PayControllerRegister()
	SupplyControllerRegister()
	err := engine.Run(":8888")
	if err != nil {
		log.Error.Println("gin启动失败", err)
	}
}
