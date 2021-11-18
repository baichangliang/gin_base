package order

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"github.com/spf13/viper"
	"net/http"
)

// UnifiedOrder 微信统一下单接口
func UnifiedOrder(c *gin.Context) {
	// 获取当前登录的用户
	//user, _ := c.Get("user")

	client := InitWxPay()
	body := "商品名"
	totalFee := "10"
	outTradeNo := "1fssdfsdfdsf0"
	openid := "os8wG5Z_ZOpTfpg7Uftnz7p5k380"
	paySign := UnifiedOrderPay(client, body, totalFee, outTradeNo, openid)
	Success(c, gin.H{"results": paySign}, "success")
}

// WxPayOrderCallback 微信回调接口
func WxPayOrderCallback(c *gin.Context) {
	notifyReq, _ := wechat.ParseNotifyToBodyMap(c.Request)
	apiKey := viper.GetString("wxPay.apiKey")

	ok, _ := wechat.VerifySign(apiKey, wechat.SignType_MD5, notifyReq)
	rsp := new(wechat.NotifyResponse)
	// ==异步通知，返回给微信平台的信息==
	if ok {
		rsp.ReturnCode = gopay.SUCCESS
		rsp.ReturnMsg = gopay.OK
		// 项目逻辑处理
	} else {
		rsp.ReturnCode = gopay.FAIL
		rsp.ReturnMsg = gopay.FAIL
	}
	c.String(http.StatusOK, "%s", rsp.ToXmlString())
}

// WxRefundOrder 微信退款申请
func WxRefundOrder(c *gin.Context) {
	client := InitWxPay()
	totalFee := "10"
	refundFee := "1"
	outTradeNo := "1fssdfsdfdsf0"

	err := UnifiedOrderRefund(client, outTradeNo, totalFee, refundFee)
	if err != nil {
		Fail(c, gin.H{}, "fail")
	}
	Success(c, gin.H{}, "success")
}
