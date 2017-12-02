package WePay

import (
	"errors"
)

type WebRequest struct {
	AppID     string `xml:"appid";json:"appid"`
	TimeStamp string `xml:"timeStamp";json:"timeStamp"`
	NonceStr  string `xml:"nonceStr";json:"nonceStr"`
	Package   string `xml:"package";json:"package"`
	SignType  string `xml:"signType";json:"signType"`
	PaySign   string `xml:"paySign";json:"paySign"`
}

func (s *WebRequest) ToMap() map[string]string {
	mapData := map[string]string{
		"appId":    s.AppID,
		"nonceStr": s.NonceStr,
		"package":  s.Package,
		"signType": s.SignType,
		"paySign":  s.PaySign,
	}

	return mapData
}

func (u *WebRequest) SignUp(key string) (err error) {
	if key == "" {
		return errors.New("key is nil")
	}
	if u == nil {
		return errors.New("WePay is nil")
	}
	m := u.ToMap()
	u.PaySign, err = getSignStr(m, key)
	if err != nil {
		return err
	}
	return nil
}
