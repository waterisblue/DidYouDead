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
	kPrivateKey = "MIIEpQIBAAKCAQEAndMBdmymPlmAe1Z5dwc/Jnz5uGEaYBpbgixq4kZ759qJk5VNWeS/0mm/aJbOKHUWJieOW9yZNBgSPIBy+4bhU3h6mDh/C/XpoBXQ2kg0Cs/BWro43bcLhIi/ICVLZogKHbNVX3/DS35isMoAgRgm7XV2bRdO5mdHKJzbHCq/9+LvCq4YMJqwpBPSUj9Q++Fj/2ikSjefk06T/qiHynNhUpC7arF0ALalGJ4F9FXhrs4vw5lFT/POUAmWHrpxVvnFa477IvhXlZGGpZrofTnXKHnXST5JgUETIfvDGtQmZGoPTTuPyT0LOVhK0p+UsSKKVNhKZLjJHCszS/syOgnDuwIDAQABAoIBAQCJI1frAn/2P6lgAwUqw5ZBKGmJqH7VsTGAzH3Xvu1dy0yufBapnRZ9jLvw8tGS4CLt5oNSGT7awGvgPLEtFMfo0xCcAsOKrhVisHI5qnHmTx6xao2NQsf0UyrYkO+S2XUI19LAV/FoQRsaiXyksErrHZ1ZL960YFu504qlUBRdHWpLcVLqHpMhxW71YDNeY6XgByLkI2TEOQgrBAjl2yusppx/OdBo/oQRdGMpcC6RHq8D3ASKOXkHBJqzbvpD4iY82VKcD274Eg3EhZ1PZASgVG/rHkwefBxb+B+QwIa1vjGr6zf93FtXkHQD00dv2DQiKPqkZagoWccahKGjzeHhAoGBANyKpXcqUDrjUaSqcgAVqGIHi6VewPqL6vJFoAqM9qVd2A2yY40m32f/MIo+dCvYMIolnWaA1+QiFCuTOQSINstvXztlmlA7C5MHGf7+XCyHyf/0R2oRh4UD0jP3P0yD84vscWAy8MTvtpLae6CNuVsObIrHwxXMVdE2FKoL6oILAoGBALcy8+GPBDj7g/RBJOfbYTMvdmEkqWg/iIKjBvhcZJLQlN87rBYxZMfi3WcTzFbQJCviLv2Gaer+2AyIRTOPLBZsTdETqA3yM+Sn5bIFr6o4tllwoxmMrubX5HeAsJ8jex1YZbkQuEaclPeQiJBwSLJKx+sxxRkcPpib6OnvcwMRAoGBAJfJzQxxUPrWzbBZp1Crl1KxAsn0pzlJ55CLnxTj5DX4muUCbEYQBR8cohnj1iz/BJlwPh74ep6HXf6Hu2yFG8F0i8TjrwudN+AfuSKpimMTqlHBnQ1Bt1YzRQ+DEkvlXVBmkbm2FcRVgAEe4oukP3iOOuJhjsGH+2CcqihHP0X1AoGBAIOs0tA/3Fgcp7ZrEFy6OP3yUlL2Z/EvIS5MgVkIB53i3+4n1MPK0tY9AMfy2f9X8Xe9ReshroWikKVJWPv7H0Qqqi/pwfPsL1JdP2KaZa2iv6Y/qGtXy+rOKD+gttJbTUVFYTAIjXrrVsMuFopCum20BEfPaQxoWbGU9m8WlYMRAoGAfpPZa3QWjlOqF3qJuIRGAbeczAFRuuGVaSBhaAe1P20Faeune/9hOKVAj1ZD+t/ff8dq/HYlK3pdcTJipmMAfhZlwQ7C+kxe/16GdpA5hFC63XEnYgdzsVYBWNUEgQs82qUEhQmSzMGkofQSm2u2g+EQ8TNImLw3e/w5V0bXDU4="
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
