package main

import (
	"dyd/mysqlconn"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	DB, err := mysqlconn.Open()

	err = DB.Ping()

	if err != nil {
		fmt.Println("数据库连接失败")
		fmt.Print(err)
		return
	}
	fmt.Print("连接成功")
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			// c.JSON：返回 JSON 格式的数据
			"message": "Hello world!",
		})
	})
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run()
}
