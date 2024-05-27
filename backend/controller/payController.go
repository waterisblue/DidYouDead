package controller

import (
	"dyd/pay"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayControllerRegister(handler ...func() func(*gin.Context)) {
	engine := getEngine()
	engine.GET("/pay", payController)
	engine.GET("/alipay/callback", pay.Callback)
	engine.GET("/alipay/notify", pay.Notify)
}

func payController(c *gin.Context) {
	url := pay.Pay()
	c.Redirect(http.StatusTemporaryRedirect, url.String())
}
