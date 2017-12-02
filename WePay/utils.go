package WePay

import (
	"errors"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"bytes"
	"strconv"
	"time"
)

//获取随机字符串
func createNonceStr()string{
	md5fun := md5.New()
	signature := make([]byte, md5fun.Size()*2)
	md5fun.Write([]byte(strconv.FormatInt(time.Now().UnixNano(),10)))
	hex.Encode(signature, md5fun.Sum(nil))
	return string(bytes.ToUpper(signature))
}

//获取签名
func getSignStr(m map[string]string, key string) (sign string, err error) {
	if (key == "") {
		return "", errors.New("key is nil")
	}
	if (m == nil) {
		return "", errors.New("map is nil")
	}
	key_array := make([]string, 0, len(m))
	for k := range m {
		key_array = append(key_array, k)
	}
	sort.Strings(key_array)

	str := ""

	for _, k := range key_array {
		v := m[k]
		if v == "" {
			continue
		}
		str += k
		str += "="
		str += v
		str += "&"
	}
	str += "key="
	str += key
	md5fun := md5.New()
	signature := make([]byte, md5fun.Size()*2)
	md5fun.Write([]byte(str))
	hex.Encode(signature, md5fun.Sum(nil))
	return string(bytes.ToUpper(signature)), nil
}

//创建订单号
func CreateOutTradeNO()string{
	return strconv.FormatInt(time.Now().UnixNano(),10)
}