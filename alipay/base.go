package alipay

import (
	"time"
	"encoding/json"
	"net/url"
	"fmt"
	"strings"
)

type (
	APC struct {
		GateWay           string
		AppId             string
		AppPrivateKey     string
		Format            string
		Charset           string
		AliPayPublicKey   string
		SignType          string
		NotifyUrl         string
		SellEmail         string
		SellId            string
	}
)

const (                                         //todo method
	PayMethod          = "alipay.trade.app.pay"
	CloseMethod        = "alipay.trade.close"
	QueryMethod        = "alipay.trade.query"
	RefundMethod       = "alipay.trade.refund"
	RefundQueryMethod  = "alipay.trade.fastpay.refund.query"
	BillDownLoadMethod = "alipay.data.dataservice.bill.downloadurl.query"

	PRODUCT_CODE       = "QUICK_MSECURITY_PAY"
)


func signStringFront(appId string) string{
	return "app_id="+appId
}

func (a *APC) signStringBack(method string) (string,string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	utfStr := "&charset="+a.Charset+"&method="+method+"&notify_url="+a.NotifyUrl+"&sign_type="+a.SignType+"&timestamp="+timeStr+"&version=1.0"
	urlStr := "&charset="+a.Charset+"&method="+method+"&notify_url="+url.QueryEscape(a.NotifyUrl)+"&sign_type="+a.SignType+"&timestamp="+AddChange20(timeStr)+"&version=1.0"
        return utfStr,urlStr
}

func (this *APC) SignStringAll(t *Pay) (string,string) {
        var middle string
	jsonByte,_ := json.Marshal(t.BizContent)
	middle = string(jsonByte)
	utfStrBack,urlStrBack := this.signStringBack(t.Method)
	utfStr := signStringFront(this.AppId)+"&biz_content="+middle+ utfStrBack
	urlStr :=  signStringFront(this.AppId)+"&biz_content="+AddChange20(middle)+ urlStrBack
        return utfStr,urlStr
}

func AddChange20(s string) string {
	return strings.Replace(url.QueryEscape(s),"+","%20",-1)
}

func (this *APC) SHA1RSASign(t *Pay) (string,error){
	ch := make(chan int,1)
	var data,urlData string
	go func() {
		data,urlData = this.SignStringAll(t)
		ch <- 1
	}()
	privateKey := createPrivateKey(this.AppPrivateKey)
	    <- ch
	if signByte,err := RSA1Sign(data,privateKey);err != nil{
		return "SHA1RSASign get failed!",err
	}else{
		return urlData+"&sign="+url.QueryEscape(Base64(signByte)),nil
	}
}

func (this *APC) SHA256RSASign(t *Pay) (string,error){
	ch := make(chan int,1)
	var data,urlData string
	go func() {
		data,urlData = this.SignStringAll(t)
		ch <- 1
	}()
	privateKey := createPrivateKey(this.AppPrivateKey)
	<- ch
	if signByte,err := RSA2Sign(data,privateKey);err != nil{
		return "SHA256RSASign get failed!",err
	}else{
		return urlData+"&sign="+url.QueryEscape(Base64(signByte)),nil
	}
}

func (this *APC) SHA1RSAVerify(data,sign string) (string,error){

	publicKey := createPublicKey(this.AliPayPublicKey)

	if err := RSA1Verify(data,sign,publicKey);err != nil{
		return "SHA256RSAVerify get failed!",err
	}else{
		return "SHA256RSAVerify get success!",nil
	}
}

func (this *APC) SHA256RSAVerify(data,sign string) (string,error){

	publicKey := createPublicKey(this.AliPayPublicKey)

	if err := RSA2Verify(data,sign,publicKey);err != nil{
		return "SHA256RSAVerify get failed!",err
	}else{
		return "SHA256RSAVerify get success!",nil
	}
}




