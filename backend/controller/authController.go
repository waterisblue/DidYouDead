package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	engine.POST("/grantAuth", GrantAuth)
}

func GrantAuth(c *gin.Context) {
	var user entity.UserInfo

	err := c.ShouldBind(&user)

	log.Info.Println("用户：", user.Username, "开始授权")
	if err != nil {
		log.Warning.Println("/grantAuth 解析失败", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 数据库校验
	exists, userdetail := dao.GetUserByAccount(user.Username)
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "该账号不存在！",
		})
		return
	}

	if userdetail.Password != user.Password {
		log.Info.Println("用户密码输入错误，输入的密码：", user.Password, "正确的密码：", userdetail.Password)
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "密码错误，请重新输入！",
		})
		return
	}

	tokenString, _ := utils.EnruptUserToken(user.Username, user.Password)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "鉴权成功",
		"data": gin.H{"token": tokenString},
	})
}
