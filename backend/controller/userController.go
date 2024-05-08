package controller

import (
	"dyd/entity"
	"dyd/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	engine.GET("/registerUser", registerUser)
}

func registerUser(c *gin.Context) {
	var user entity.UserInfo

	err := c.ShouldBind(&user)
	log.Info.Println("用户注册，用户名：", user.Username, "用户密码：", user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误的参数",
		})
		return
	}

	// 插入数据库

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
		"data": user,
	})
}
