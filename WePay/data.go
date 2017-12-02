package WePay

type CDATA struct {
	Text string `xml:",cdata"`
}

const (
	UNIFIEDORDER_URL="https://api.mch.weixin.qq.com/pay/unifiedorder"
	ORDERQUERY_URL="https://api.mch.weixin.qq.com/pay/orderquery"
)