package pay

import (
	"dyd/dao"
	"dyd/entity"
	"dyd/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"net/http"
	"net/url"
)

const (
	kAppId = "9021000137647008"
	//私钥
	kPrivateKey = "MIIEpQIBAAKCAQEAoFy2NjoeJ3yR5ANrWAoRMmqE5KrUWBt7FmvwAXWbV62lDGm7RQ0mvXMpJGTMPzIIdHJz1+xPqur+grgWP/m6j6dRNomr0tDfb3CP/1UVphFNPxJoQsA03h3jcuG9xbb5uk7dEd2kwRbR1ZThZX7rj/2RfCcS2Pf+5KGQ2r6moRrOriyQeIpulCB+TZHCFihzjdGKsRm7IyY/v8+PcjYbMefJrb8P7T5QbSTNZ0rqULK4BdklFeAPF2MjklSYL2KsRqVtrqONtM5USYeiq3dGd2HcouKeJlgbw3sQAUhuSvyH/TU8192TOqA12IgUiPZe+2MnnX01EAlaBKpVENEgGQIDAQABAoIBAQCbTK35nhcNzonK87TFMhqtZJMGkJnMKfDuKeHox+iE5NluSAbgPHhECH7Ti/pxSRb/Dsg5V9VfhOyt03FYy5RtAnCYDIOmSjnYRn5tnAy3IxJX/o+R9Wo9oVxgKVYLX8RBNZbKcofXpy3XVbNjA6NFcx+tyHlsPd9Ps1gBP5lHcolBE/O8HOudFCymCjEBU6GsNrai+O6j+Ss8kxz89LEwPDFzha5b/kL/uUJOhpop5bm/RsxwvUgdj6QZjMypbFJKgaPn+Fk034xq6pcQ3rO7YsQDGftnpRi1n6QohftJl4zUyLldBtfwkozvh0AzldHE9gwUYmFSTAXzotk7XOfBAoGBAPraXMRqHCAMkVN3+nOZkZ0v8CgT4XIGI2ENLcKYBcZb8EYJNsPdS8ZSD7bpoM62sk82BwOpVF7m6RDLE6GrxDsHWO6r6r5CQaiU0keIcvfUGrGHiUMYhoGWQKtifeulyN8XgcPoc2VvR5bOcblFypn/puDJeA98doSWKkstXoBNAoGBAKOnCOBWMVppjPdTR80wmE9CoNp3udYGaHg5FZKC1Pisi4WUd2GHsTuOkb96/6oMMwYI4o5lS2PtlfZo5dg80YYtjTNJZ4m+P2pWLsHZRtukPGjL0OIbKN7Iui8msjqwWuugajcBEIzLWi4f3h0y+vVTeTgWg9dSweVbfa32p6T9AoGBAMLR6A4bIrC5UfxzR62T7QUY//kAR6zYm416QLnKe2GscZTS/xdqvLjPbqjuFRe1yBHAsf3j+kSiLQv/y8VZs1Lsx+LE3dDZgEi/G/i2PiO/NOP/kVjsqTnnSV5RZd4nqTOp+/v8gfKbVAwJVUVAsLKdXCtkDLxqO1mJjltIZwZZAoGAbuLhGtCMREy8nWCqiBfjO+FTiHQTUTeqVbuoKEvJ7m1LZEmC0mzU5BvoGaxHy3rUn8Qpbn6oJJ95Oys8gdZDyfzXbebQ5v4nje+zYtnML0sLO57OfEYW3U15CjXYEfgowvtAvOBVzHCFPDePko1Wih/zRNPMpXuHsdbaQ6ObWCkCgYEAqpS295a/BClg+ZS2ZyOtlegZ9jvxXXzwZwIWoxjrNseC76bXMNVRVyMfCGkx+lxOMQKgoFLA5XjTFlOAImoA9Chg0ZXlO1mJVQ/CV0jTLJCs8aG578Z61wUt69FOHUTI3bXKN8ZZoH7R2YKmGNck1ktdSJdZ1qkJ+0WycKpUiMA="
	// TODO 设置回调地址域名
	kServerDomain = ""
)

var Client *alipay.Client
var err error

func init() {
	if Client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Warning.Println("初始化支付宝失败", err)
		return
	}
	if err = Client.LoadAppCertPublicKeyFromFile("./certification/appPublicCert.crt"); err != nil {
		log.Warning.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝根证书
	if err = Client.LoadAliPayRootCertFromFile("./certification/alipayRootCert.crt"); err != nil {
		log.Warning.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝公钥证书
	if err = Client.LoadAlipayCertPublicKeyFromFile("./certification/alipayPublicCert.crt"); err != nil {
		log.Warning.Println("加载证书发生错误", err)
		return
	}
	//接口内容加密
	//	if err = Client.SetEncryptKey("iotxR/d99T9Awom/UaSqiQ=="); err != nil {
	//log.Warning.Println("加载内容加密密钥发生错误", err)
	return
}

func Pay(amount string, subject string, typeId int, id int) *url.URL {
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = fmt.Sprintf("%s:%s", subject, tradeNo)
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	urlRes, _ := Client.TradePagePay(p)
	orderEntity := entity.OrderEntity{
		OrderNum:  tradeNo,
		OrderType: typeId,
		LinkId:    id,
	}
	go dao.InsertPayDao(orderEntity)
	return urlRes
	//http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
}

func Callback(c *gin.Context) {
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Warning.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}

	// 获取通知参数
	outTradeNo := c.Request.Form.Get("out_trade_no")

	// 验证签名
	if err := Client.VerifySign(c.Request.Form); err != nil {
		log.Warning.Println("回调验证签名发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "回调验证签名发生错误"})
		return
	}

	log.Warning.Println("回调验证签名通过")

	// 查询订单状态
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := Client.TradeQuery(c, p)
	if err != nil {
		log.Warning.Printf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Warning.Printf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}
	go dao.UpdatePayDao(true, outTradeNo)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("订单 %s 支付成功", outTradeNo)})
}
func Notify(c *gin.Context) {
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Warning.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}

	// 解析异步通知
	notification, err := Client.DecodeNotification(c.Request.Form)
	if err != nil {
		log.Warning.Println("解析异步通知发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析异步通知发生错误"})
		return
	}

	log.Warning.Println("解析异步通知成功:", notification.NotifyId)

	// 查询订单状态
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err := Client.Request(c, p, &rsp); err != nil {
		log.Warning.Printf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Warning.Printf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}
	go dao.UpdatePayDao(true, notification.OutTradeNo)
	log.Warning.Printf("订单 %s 支付成功", notification.OutTradeNo)

	Client.ACKNotification(c.Writer)
}
