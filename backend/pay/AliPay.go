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
	kPrivateKey = "MIIEowIBAAKCAQEAk/6vrImbRh5XJb2InqZiPV8oQdm8bMs+p/RAkHJl3UBb5q8LIirQvwGQ07AiKh9slV81JGxjDVfOLXovdDY1x7RFZKqrSIytVHOoU1ugkfndgj0DHbl6bFjhMjDQ6JzhGE56aH71DOMoMGIenadYaYPb3AmthNQZfu2KgohxsF3lipQuL1LjaTlhbI2q3n2gdeQQthhixxzkD9xArKWE1r7Dy1UZsAjjrHdF8MG74KV4cMIuvVBU/0nl1uHxGjs7qBqoCIVk38ODjqyPyDycNDom/ageT/byMFxKnVSEkn7Ynkev90vJjnyRVJ/PCv3Ry9/4R8gt9S5JtB0/zF75vQIDAQABAoIBABWq1ReCsgwLByFMZYMswvdPRjqV4UNgYlT+qrE7PnStJodYm12uPQ+p8AU+JVJbdW9NGB+0CrB6aYI5AFeYVDlfyEUbw6YEip/xqvRBVvnoRij6O8mPTUAtpxLNGNjllwMTAlLoO4Jeg5TJ2MemJ4iTOI43kIMA9pahuoSv600xyRKj1oFBpw08yyciMrr4XMU5TLpjZGqVJ2zbkMim+WsmV8huAYAZmvdQHK2OyEcu0de6qYBbz/oSuacGxOyG9migIEhnCKQyJhAfZ+Q34DRsd1ahtJRGHXL/HUrECdmsgtaT3mehoPL11VgJhc4ZOGj3YIHW+CdcmuXRVvon5pUCgYEAyNICjXbhWY/iK6EWAqanwbIIXubqW4MSL2+GNT4aivYZqS7NoeMUkktKKkGdtEqtUJSLzA5fod66732XFUURlPNg0cynf4hv1lFEYHtLMMKm6Y/sEbNy0A7U5amlwC2zM0HdT0PSRi7KqEQZwqPxYjtc9q/cQsIm+YT1VXwimsMCgYEAvKjbPFGqcJyzKuk7yh/mhpjLO+O4eqIaLhpEO0p6KlpQjhdKxYSshMuTgiV07fJcjdGUgiOh9NgumXQmA6HVE1Vhyx7iHDeHe8RRbQXRocHdobfg8V7/uZ7HcllGALnll6vnneSXug2hoIzbzdHuI8lmf64UlIFef5WThgIJ0X8CgYEAs2eKiD7QC2cTMs7yDuQ7sjsOZ1n7vE1YDbCgQSh+Je699fquEhX/5sIEO8Pq04pzUy2hxmaA4OdOMW4Nbx8JGfxjXopqiPCmmNx5LYh4H9OFsdq1lwqtY1ocsB/ix9EL6prIz2tTiwd6XOCZ8tzrZEqMPoXVTH6OuxMg6YHAmGkCgYBN/S/EG65tpxmZ35jL0awmb/tz4otchq5z6H/UUtF1uKgcVRyf0lzO+Kkd4AYaZ1pDdiQCOGC3Krb40L72OWwv9dfbdATczUfArkNK4mDqkY3hUHY40aC9RxoD90us+SwkKvgwmVLvorrx98jVqLbdLGCm33u6eFBsRZUajLUqWQKBgE3onCcLTjFKXNrQCJ6CcTocZG5SmiJIXJtusLioIfsVxbFEL6cljUXKPzcW/WVAQ456L5+SRSwhp1b1GlQ0eEgNx++teW5ypF0SB93+VKUKA3UO2/BCKhzWvFGC6aZuKLBux5J0n4CS0ZLUCcXLF+rw+Rg5juB2hPIWb3b/StN0"
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
