package main

import (
	"dyd/controller"
	"dyd/mysqlconn"
	"dyd/udpServer"
)

func main() {
	// 初始化数据库连接
	mysqlconn.Open()
	// 启动UDP
	go udpServer.Start()
	go udpServer.RandomWrite()
	// 启动gin
	controller.StartGin()
}
