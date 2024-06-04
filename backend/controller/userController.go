package controller

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"dyd/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()

	engine.POST("/registeruser", registerUser)
	engine.POST("/getuseradmin", middleware.JWTAuthMiddleware(), getUserAdministartor)

	loginafter := engine.Group("/loginafter", middleware.CorsMiddleware(), middleware.JWTAuthMiddleware())
	loginafter.POST("/getalluser", GetAllUser)
	loginafter.POST("/deleteuser", DeleteUserByAccount)
	loginafter.POST("/grant", GrantUser)
}

// 注册
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

	// 判断账号是否存在
	exists, _ := dao.GetUserByAccount(user.Username)
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "该账号已经存在，请重新填写",
		})
		return
	}
	// 插入数据库
	dao.AddUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
		"data": user,
	})
}

// 获取用户权限
func getUserAdministartor(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "用户未登录！",
		})
		return
	}
	log.Info.Println(username, "用户开始获取权限，访问了", "getUserAdministartor")

	_, userdetail := dao.GetUserByAccount(username.(string))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "权限获取成功！",
		"data": gin.H{"admin": userdetail.Administrator},
	})
}

func GetAllUser(c *gin.Context) {
	_, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "用户未登录",
		})
		return
	}
	userLists := dao.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户成功",
		"data": userLists,
	})
}

func DeleteUserByAccount(c *gin.Context) {
	var user entity.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数校验失败",
		})
	}
	log.Info.Println(user.Username)
	go dao.RemoveUserByAccount(user.Username)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GrantUser(c *gin.Context) {
	var user entity.UserDetail
	c.ShouldBind(&user)

	go dao.GrantUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}
