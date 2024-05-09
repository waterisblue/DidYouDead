package main

import (
	"dyd/controller"
	log "dyd/log"
	"dyd/mysqlconn"
)

func main() {
	// 初始化日志
	log.Init()

	// 初始化数据库连接
	mysqlconn.Open()

	// 初始化gin
	controller.InitGin()
	controller.StartGin()

}
