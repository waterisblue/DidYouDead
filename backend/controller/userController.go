package controller

import "github.com/gin-gonic/gin"

func UserControllerRegister() {
	engine := getEngine()
	engine.GET("/users", GetUsers)
}

func GetUsers(c *gin.Context) {
	res := "users"
	c.String(200, res)
}
