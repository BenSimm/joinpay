package joinpay

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"github.com/parnurzeal/gorequest"
)

type joinClient struct {
	MerchantNo 	string
	ApiKey 		string
	isProd 		bool
}

/**
初始化汇聚客户端
MerchantNo	汇聚商户编号
 */
func New(MerchantNo string , Apikey string, isProd bool) (client *joinClient) {
	client = new(joinClient)
	client.MerchantNo = MerchantNo
	client.ApiKey = Apikey
	client.isProd = isProd
	return client
}

/**
汇聚统一下单
 */
func (this *joinClient) UniPayApi(body BodyMap) (joinRsp *UniPayApiResponse, err error) {
	var bytes []byte
	if !this.isProd {
		body.Set("total_fee", 101)
	}

	tlsConfig := new(tls.Config)
	tlsConfig.InsecureSkipVerify = true

	bytes, err = this.doJoin(body, joinURL_UniPayApi, tlsConfig)
	if err != nil {
		return nil, err
	}

	joinRsp = new(UniPayApiResponse)
	err = json.Unmarshal(bytes, joinRsp)
	if err != nil {
		return nil, err
	}
	return joinRsp, nil
}

/**
汇聚退款
 */
func (this *joinClient) Refund(body BodyMap) (joinRsp *RefundResponse , err error){
	var bytes []byte

	tlsConfig := new(tls.Config)
	tlsConfig.InsecureSkipVerify = true

	bytes, err = this.doJoin(body, joinURL_Refund, tlsConfig)
	if err != nil {
		return nil, err
	}

	joinRsp = new(RefundResponse)
	err = xml.Unmarshal(bytes, joinRsp)
	if err != nil {
		return nil, err
	}
	return joinRsp, nil
}

//向微信发送请求 ok
func (this *joinClient) doJoin(body BodyMap, url string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	var sign string
	body.Set("p1_MerchantNo", this.MerchantNo)
	//===============生成参数===================
	//正式环境
	//本地计算Sign
	sign = getLocalSign(this.ApiKey, body.Get("sign_type"), body)

	body.Set("hmac", sign)
	//===============发起请求===================
	agent := gorequest.New()

	if this.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}

	bodyM,err := json.Marshal(body)
	bodyString := string(bodyM)

	agent.Post(url)
	agent.Type("form")
	//fmt.Printf(bodyString)
	agent.SendString(bodyString)
	_, bytes, errs := agent.EndBytes()
	//fmt.Println(string(bytes))
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return bytes, nil
}

