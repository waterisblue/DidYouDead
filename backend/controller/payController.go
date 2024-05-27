package controller

import (
	"dyd/pay"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PayControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	engine.GET("/pay", payController)
	engine.GET("/alipay/callback", pay.Callback)
	engine.GET("/alipay/notify", pay.Notify)
}

func payController(c *gin.Context) {
	amount, _ := c.GetQuery("amount")
	subject, _ := c.GetQuery("subject")
	id, _ := c.GetQuery("id")
	typeId, _ := c.GetQuery("typeId")
	idInt, _ := strconv.Atoi(id)
	typeIdInt, _ := strconv.Atoi(typeId)
	url := pay.Pay(amount, subject, typeIdInt, idInt)
	c.Redirect(http.StatusTemporaryRedirect, url.String())
}
