package WePay

import (
	"fmt"
	"errors"
)

type UnifiedOrderRequest struct {
	AppID          string `xml:"appid"`            //公众账号ID
	MchID          string `xml:"mch_id"`           //商户号
	DeviceInfo     string `xml:"device_info"`      //设备号
	NonceStr       string `xml:"nonce_str"`        //随机字符串
	Sign           string `xml:"sign"`             //签名
	SignType       string `xml:"sign_type"`        //签名类型
	Body           string `xml:"body"`             //商品描述
	Attach         string `xml:"attach"`           //附加数据
	OutTradeNo     string `xml:"out_trade_no"`     //商户订单号
	FeeType        string `xml:"fee_type"`         //标价币种
	TotalFee       int    `xml:"total_fee"`        //标价金额
	SpbillCreateIP string `xml:"spbill_create_ip"` //终端IP
	TimeStart      string `xml:"time_start"`       //交易起始时间
	TimeExpire     string `xml:"time_expire"`      //交易结束时间
	NotifyUrl      string `xml:"notify_url"`       //通知地址
	TradeType      string `xml:"trade_type"`       //交易类型
	OpenID         string `xml:"openid"`           //用户标识
}


func (s *UnifiedOrderRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"appid":            s.AppID,
		"mch_id":           s.MchID,
		"device_info":      s.DeviceInfo,
		"nonce_str":        s.NonceStr,
		"body":             s.Body,
		"attach":           s.Attach,
		"out_trade_no":     s.OutTradeNo,
		"fee_type":         s.FeeType,
		"total_fee":        fmt.Sprintf("%d", s.TotalFee),
		"spbill_create_ip": s.SpbillCreateIP,
		"time_start":       s.TimeStart,
		"time_expire":      s.TimeExpire,
		"notify_url":       s.NotifyUrl,
		"trade_type":       s.TradeType,
		"sign_type":        s.SignType,
		"openid":           s.OpenID,
	}

	return mapData
}

func (u *UnifiedOrderRequest) SignUp(key string)(err error) {
	if(key==""){
		return errors.New("key is nil")
	}
	if(u==nil){
		return errors.New("WePay is nil")
	}
	m := u.ToMap()
	u.Sign,err=getSignStr(m,key)
	if(err!=nil){
		return err
	}
	return nil
}
