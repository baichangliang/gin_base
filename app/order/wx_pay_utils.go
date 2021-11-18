package order

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

// InitWxPay 微信初始化
func InitWxPay() *wechat.Client {
	appid := viper.GetString("wxPay.appid")
	mchId := viper.GetString("wxPay.mchId")
	apiKey := viper.GetString("wxPay.apiKey")

	client := wechat.NewClient(appid, mchId, apiKey, true)
	client.DebugSwitch = gopay.DebugOn
	//  微信证书二选一：只传 apiclient_cert.pem 和 apiclient_key.pem 或者只传 apiclient_cert.p12
	//err := client.AddCertPemFilePath("./cert/apiclient_cert.pem", "./cert/apiclient_key.pem")
	//if err != nil {
	//	return nil
	//}
	return client
}

// UnifiedOrderPay 微信统一下单接口
func UnifiedOrderPay(cli *wechat.Client, body, totalFee, outTradeNo, openid string) (ma interface{}) {

	appid := viper.GetString("wxPay.appid")
	apiKey := viper.GetString("wxPay.apiKey")

	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", RandomString(32))
	bm.Set("body", body)
	bm.Set("notify_url", "https://order.lzzhuangyuan.cn:9879/api/order/callback")
	bm.Set("trade_type", "JSAPI")
	bm.Set("spbill_create_ip", "125.75.47.3")
	bm.Set("total_fee", totalFee)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("openid", openid)

	// 我当前使用的版本是1.5.53--如果你想使用最新的那么这里可能需要两个参数
	wxRsp, _ := cli.UnifiedOrder(bm)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	prepayId := "prepay_id=" + wxRsp.PrepayId
	paySign := wechat.GetMiniPaySign(appid, wxRsp.NonceStr, prepayId, wechat.SignType_MD5, timeStamp, apiKey)
	res := make(map[string]string)
	res["package"] = prepayId
	res["timeStamp"] = timeStamp
	res["nonceStr"] = wxRsp.NonceStr
	res["appId"] = wxRsp.Appid
	res["signType"] = "MD5"
	res["paySign"] = paySign
	res["paySign"] = paySign
	return res
}

// UnifiedOrderRefund 微信退款申请
func UnifiedOrderRefund(cli *wechat.Client, outTradeNo, totalFee, refundFee string) (ma interface{}) {

	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", RandomString(32))
	bm.Set("outTradeNo", outTradeNo)
	bm.Set("outTradeNo", totalFee)
	bm.Set("outTradeNo", refundFee)
	_, _, err := cli.Refund(bm)
	return err
}

// RandomString 获取32位随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
