package WePay

import (
	"errors"
)

type QueryRequest struct {
	AppID          string `xml:"appid"`            //公众账号ID
	MchID          string `xml:"mch_id"`           //商户号
	OutTradeNo     string `xml:"out_trade_no"`     //商户订单号
	NonceStr       string `xml:"nonce_str"`        //随机字符串
	Sign           string `xml:"sign"`             //签名
	SignType       string `xml:"sign_type"`        //签名类型
}

func (s *QueryRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"appid":            s.AppID,
		"mch_id":           s.MchID,
		"nonce_str":        s.NonceStr,
		"out_trade_no":     s.OutTradeNo,
		"sign_type":        s.SignType,
	}

	return mapData
}


func (u *QueryRequest) SignUp(key string)(err error) {
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
