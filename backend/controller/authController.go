package controller

import (
	"dyd/entity"
	"dyd/log"
	"dyd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	engine.GET("/grantAuth", GrantAuth)
}

func GrantAuth(c *gin.Context) {
	var user entity.UserInfo

	err := c.ShouldBind(&user)

	if err != nil {
		log.Warning.Println("/grantAuth 解析失败", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 数据库校验

	tokenString, _ := utils.EnruptUserToken(user.Username, user.Password)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "鉴权成功",
		"data": gin.H{"token": tokenString},
	})
}
