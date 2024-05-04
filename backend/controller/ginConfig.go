package controller

import (
	log "dyd/log"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

// 初始化gin
func InitGin() {
	engine = gin.Default()
}

// 获取gin
func getEngine() (engin *gin.Engine) {
	if engine == nil {
		log.Error.Println("gin engine还未进行初始化，无法获取，请先进行初始化！")
	}
	return engine
}

// 开启服务器
func StartGin() {
	UserControllerRegister()
	engine.Run()
}