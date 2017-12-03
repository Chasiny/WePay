package WePay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type WePay struct {
	UnifiedOrderUrl string
	OrderQueryUrl   string
	NotifyUrl       string
	AppId           string
	MchID           string
	Key             string
}

//初始化Wepay
func (wp *WePay) Init(appid string, mchid string, key string, notify_url string) error {
	if UNIFIEDORDER_URL == "" {
		return errors.New("UNIFIEDORDER_URL nil")
	}
	if ORDERQUERY_URL == "" {
		return errors.New("ORDERQUERY_URL nil")
	}
	if notify_url == "" {
		return errors.New("NOTIFY_URL nil")
	}
	if appid == "" {
		return errors.New("AppID nil")
	}
	if mchid == "" {
		return errors.New("MchID nil")
	}
	if key == "" {
		return errors.New("Key nil")
	}
	wp.UnifiedOrderUrl = UNIFIEDORDER_URL
	wp.OrderQueryUrl = ORDERQUERY_URL
	wp.NotifyUrl = notify_url
	wp.AppId = appid
	wp.MchID = mchid
	wp.Key = key
	return nil
}

//统一下单
func (wp *WePay) Request(out_trade_no string, fee int, client_ip string, duration int, open_id string) (*UnifiedOrderRespone, error) {
	if wp.MchID == "" || wp.AppId == "" {
		return nil, errors.New("MchID or AppId nil")
	}
	if wp.Key == "" {
		return nil, errors.New("Key nil")
	}
	if out_trade_no == "" {
		return nil, errors.New("out_trade_no nil")
	}
	randomStr := createNonceStr()
	if len(randomStr) != 32 {
		return nil, errors.New("randomStr error")
	}
	if fee <= 0 {
		return nil, errors.New("fee smaller than 0")
	}
	if duration <= 0 {
		return nil, errors.New("duration smaller than 0")
	}
	if client_ip == "" {
		return nil, errors.New("client_ip nil")
	}
	if open_id == "" {
		return nil, errors.New("open_id nil")
	}

	request := &UnifiedOrderRequest{
		AppID:          wp.AppId,
		MchID:          wp.MchID,
		DeviceInfo:     "DeviceInfo",
		NonceStr:       randomStr,
		Sign:           "Sign",
		SignType:       "MD5",
		Body:           "广州大学捐赠基金会-校友捐赠",
		Attach:         "测试",
		OutTradeNo:     out_trade_no,
		FeeType:        "CNY",
		TotalFee:       fee,
		SpbillCreateIP: client_ip,
		TimeStart:      time.Now().Format("20060102150405"),
		TimeExpire:     time.Now().Add(time.Minute * time.Duration(duration)).Format("20060102150405"),
		NotifyUrl:      wp.NotifyUrl,
		TradeType:      "JSAPI",
		OpenID:         open_id,
	}
	var err error
	err = request.SignUp(wp.Key)
	if err != nil {
		return nil, err
	}

	buf, err := post(wp.UnifiedOrderUrl, request)
	if err != nil {
		return nil, err
	}
	respon := &UnifiedOrderRespone{}
	err = xml.Unmarshal(buf, respon)
	if err != nil {
		return nil, err
	}
	return respon, nil
}

//整合获取前端下单的参数
func (wp *WePay) WebRequest(out_trade_no string, fee int, client_ip string, duration int, open_id string) (*WebRequest, error) {
	if wp.Key == "" {
		return nil, errors.New("Key nil")
	}
	res, err := wp.Request(out_trade_no, fee, client_ip, duration, open_id)
	if err != nil {
		return nil, err
	}
	if res.ReturnCode.Text != "SUCCESS" || res.ResultCode.Text != "SUCCESS" {
		return nil, errors.New("WebRequest fail:" + res.ErrCodeDes.Text)
	}
	webres := &WebRequest{
		AppID:     wp.AppId,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  createNonceStr(),
		Package:   "prepay_id=" + res.PrepayID.Text,
		SignType:  "MD5",
	}
	err = webres.SignUp(wp.Key)
	if err != nil {
		return nil, err
	}
	return webres, nil
}

//查询订单
func (wp *WePay) Query(out_trade_no string) (res *QueryRespon, err error) {
	if wp.MchID == "" || wp.AppId == "" {
		return nil, errors.New("MchID or AppId nil")
	}
	if out_trade_no == "" {
		return nil, errors.New("out_trade_no nil")
	}
	randomStr := createNonceStr()
	if len(randomStr) != 32 {
		return nil, errors.New("randomStr error")
	}

	q := &QueryRequest{
		AppID:      wp.AppId,
		MchID:      wp.MchID,
		OutTradeNo: out_trade_no,
		NonceStr:   randomStr,
		SignType:   "MD5",
	}
	err = q.SignUp(wp.Key)
	if err != nil {
		return nil, errors.New("SignUp error:" + err.Error())
	}

	respon, err := post(wp.OrderQueryUrl, q)
	if err != nil {
		return nil, err
	}
	queryRespon := &QueryRespon{}
	xml.Unmarshal(respon, &queryRespon)
	return queryRespon, nil

}

//简单封装post
func post(url string, data interface{}) (res_buf []byte, err error) {
	if url == "" {
		return nil, errors.New("url nil")
	}
	if data == nil {
		return nil, errors.New("buf nil")
	}
	buf, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	r, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(buf)))
	if err != nil {
		return nil, err
	}
	res_buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return res_buf, nil
}
