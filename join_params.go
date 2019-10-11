package joinpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

//本地通过支付参数计算Sign值
func getLocalSign(apiKey string, signType string, body BodyMap) (sign string) {
	signStr := SortJoinSignParams(apiKey, body)
	//fmt.Println("signStr:", signStr)
	var hashSign []byte
	if signType == SignType_HMAC_SHA256 {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	fmt.Println("signStr:" + signStr)
	sign = strings.ToLower(hex.EncodeToString(hashSign))
	return
}

//获取根据Key排序后的请求参数字符串
func SortJoinSignParams(apiKey string, body BodyMap) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(body.Get(k))
	}
	buffer.WriteString(apiKey)
	return buffer.String()
}


//生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")

	for key := range bm {
		buffer.WriteString("<")
		buffer.WriteString(key)
		buffer.WriteString("><![CDATA[")
		buffer.WriteString(bm.Get(key))
		buffer.WriteString("]]></")
		buffer.WriteString(key)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
