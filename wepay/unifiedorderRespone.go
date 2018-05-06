package WePay

import (
	"errors"
)

type UnifiedOrderRespone struct {
	ReturnCode CDATA `xml:"return_code"`  //返回错误码
	ReturnMsg  CDATA `xml:"return_msg"`   //返回信息
	AppID      CDATA `xml:"appid"`        //公众账号ID
	MchID      CDATA `xml:"mch_id"`       //商户号
	DeviceInfo CDATA `xml:"device_info"`  //设备号
	NonceStr   CDATA `xml:"nonce_str"`    //随机字符串
	Sign       CDATA `xml:"sign"`         //签名
	ResultCode CDATA `xml:"result_code"`  //业务结果
	ErrCode    CDATA `xml:"err_code"`     //错误代码
	ErrCodeDes CDATA `xml:"err_code_des"` //错误代码描述
	TradeType  CDATA `xml:"trade_type"`   //交易类型
	PrepayID   CDATA `xml:"prepay_id"`    //预支付交易会话标识
}

func (s *UnifiedOrderRespone) ToMap() map[string]string {
	mapData := map[string]string{
		"return_code": s.ReturnCode.Text,
		"return_msg":  s.ReturnMsg.Text,
		"appid":       s.AppID.Text,
		"mch_id":      s.MchID.Text,
		"device_info": s.DeviceInfo.Text,
		"nonce_str":   s.NonceStr.Text,
		"result_code": s.ResultCode.Text,
		"trade_type":  s.TradeType.Text,
	}

	return mapData
}

func (u *UnifiedOrderRespone) SignUp(key string) (err error) {
	if (key == "") {
		return errors.New("key is nil")
	}
	if (u == nil) {
		return errors.New("UnifiedOrderRespon is nil")
	}
	m := u.ToMap()
	u.Sign.Text,err=getSignStr(m, key)
	if(err!=nil){
		return err
	}
	return nil
}
