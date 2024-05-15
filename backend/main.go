package main

import (
	"dyd/controller"
	"dyd/mysqlconn"
)

func main() {
	// 初始化数据库连接
	mysqlconn.Open()
	// 启动gin
	controller.StartGin()
}
