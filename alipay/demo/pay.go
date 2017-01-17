package alipay

import (
	"github.com/maomaoxgg/third_party_pay/alipay"
	"strconv"
)

type Order struct {
	Id 		int
	Amount		int
	Tid		string                //服务端自己生成的订单号，放进支付宝的OutTradeNo中，业务处理成功后，支付宝会根据你这个订单号生成一个对应你的appid的支付宝内部TradeNo，大致逻辑是这样
}

const (
	GATE_WAY                    = "https://openapi.alipay.com/gateway.do"
	SAND_GATE_WAY               = "https://openapi.alipaydev.com/gateway.do"
	APP_ID                      = "2017777777777777"
	SAND_BOX_ID                 = "2016777777777777"

	//key写在代码中注意每行64字符，如果从文件载入不存在此问题
	APP_PRIVATE_KEY             = `
-----BEGIN RSA PRIVATE KEY-----
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
-----END RSA PRIVATE KEY-----
`

	/*
	todo 支付宝公钥
	 */
	ALI_PAY_PUBLIC_KEY          = `
-----BEGIN RSA PUBLIC KEY-----
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
-----END RSA PUBLIC KEY-----
`
	ALI_CHARSET                 = "UTF-8"
	FORMAT                      = "json"
	SIGN_TYPE                   = "RSA2"
	NOTIFY_URL                  ="https://xxx.xx.xxx.xxx/xxx/xxxxx"
	SELL_EMAIL                  ="golang@163.com"    //自用
)


var A *alipay.APC

func init(){
	A = &alipay.APC{GateWay:GATE_WAY,Charset:ALI_CHARSET,Format:FORMAT,AppId:APP_ID,AliPayPublicKey:ALI_PAY_PUBLIC_KEY,AppPrivateKey:APP_PRIVATE_KEY,SignType:SIGN_TYPE,NotifyUrl:NOTIFY_URL,SellEmail:SELL_EMAIL}
}

func getSign(t *alipay.Pay,mode int) string {
	if mode == 1{
		A.GateWay = SAND_GATE_WAY;A.AppId = SAND_BOX_ID   //沙盒
	}

	s,err := A.SHA256RSASign(t)
	if  err != nil{
		fmt.Println(err)
	}
	return s
}

func (this *Order)CreateAliPayRechargeOrder(title,note string) interface{} {
	//todo 这里写自己的订单逻辑，也可以加上error判定，我把自己的都删了，只留下关键部分
			TA := strconv.FormatFloat(float64(this.Amount)/100,'f',-1,64)
			b := &alipay.Biz{OutTradeNo:this.Tid,TimeoutExpress:"30M",TotalAmount:TA,SellId:"",Subject:title,Body:note,ProductCode:alipay.PRODUCT_CODE}
			return AliPay(b)
}

func AliPay(b *alipay.Biz) string {
	t := &alipay.Pay{Method:alipay.PayMethod,BizContent:b}
	return getSign(t)
}
