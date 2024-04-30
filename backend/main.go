package main

import (
	controller "dyd/controller"
	log "dyd/log"
)

func main() {
	// 初始化日志
	log.Init()

	// 初始化gin
	controller.InitGin()
	controller.StartGin()
	// DB, err := mysqlconn.Open()
	return
}
