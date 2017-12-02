package WePay

import (
	"errors"
)

type WebRequest struct {
	AppID     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
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
