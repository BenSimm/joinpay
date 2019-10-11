package joinpay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"golang_payment/pkg/logging"
	"net/http"
	"net/url"
	"strings"
)

var (
	kSuccess = []byte("success")
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
	body.Remove("sign_type")
	//===============发起请求===================
	agent := gorequest.New()

	if this.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}

	bodyM,err := json.Marshal(body)
	bodyString := string(bodyM)

	agent.Post(url)
	agent.Type("form")
	agent.SendString(bodyString)
	_, bytes, errs := agent.EndBytes()
	//fmt.Println(string(bytes))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	fmt.Println(string(bytes) + "123123123213")

	return bytes, nil
}

//支付通知的签名验证和参数签名后的Sign
//    apiKey：API秘钥值
//    signType：签名类型 MD5 或 HMAC-SHA256（默认请填写 MD5）
//    notifyRsp：利用 gopay.ParseNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非微信返回参数中的Sign
func VerifyPayResultSign(apiKey string, signType string, notifyRsp *JoinNotifyRequest) (ok bool, sign string) {

	body := make(BodyMap)
	body.Set("r1_MerchantNo", notifyRsp.R1MerchantNo)
	body.Set("r2_OrderNo", notifyRsp.R2OrderNo)
	body.Set("r3_Amount", notifyRsp.R3Amount)
	body.Set("r4_Cur", notifyRsp.R4Cur)
	body.Set("r5_Mp", notifyRsp.R5Mp)
	body.Set("r6_Status", notifyRsp.R6Status)
	body.Set("r7_TrxNo", notifyRsp.R7TrxNo)
	body.Set("r8_BankOrderNo", notifyRsp.R8BankOrderNo)
	body.Set("r9_BankTrxNo", notifyRsp.R9BankTrxNo)
	body.Set("ra_PayTime", notifyRsp.RaPayTime)
	body.Set("rb_DealTime", notifyRsp.RbDealTime)
	body.Set("rc_BankCode", notifyRsp.RcBankCode)

	//newBody := make(BodyMap)
	//for k, v := range body {
	//	vStr := convert2String(v)
	//	if vStr != "" && vStr != "0" {
	//		newBody.Set(k, v)
	//	}
	//}

	signStr := SortJoinSignParams(apiKey, body)
	logging.Info(signStr)
	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToLower(hex.EncodeToString(hashSign))
	ok = sign == notifyRsp.Hmac
	return
}

//解析微信支付异步通知的参数
//    req：*http.Request
//    返回参数notifyReq：Notify请求的参数
//    返回参数err：错误信息
func ParseJoinNotifyResult(req *http.Request) (notifyReq *JoinNotifyRequest, err error) {
	notifyReq = new(JoinNotifyRequest)
	err = req.ParseForm()
	notifyReq.R1MerchantNo = req.Form.Get("r1_MerchantNo")
	notifyReq.R2OrderNo = req.Form.Get("r2_OrderNo")
	notifyReq.R3Amount = req.Form.Get("r3_Amount")
	notifyReq.R4Cur = req.Form.Get("r4_Cur")
	notifyReq.R5Mp = req.Form.Get("r5_Mp")
	notifyReq.R6Status = req.Form.Get("r6_Status")
	notifyReq.R7TrxNo = req.Form.Get("r7_TrxNo")
	notifyReq.R8BankOrderNo = req.Form.Get("r8_BankOrderNo")
	notifyReq.R9BankTrxNo = req.Form.Get("r9_BankTrxNo")
	notifyReq.RaPayTime,_ = url.QueryUnescape(req.Form.Get("ra_PayTime"))
	notifyReq.RbDealTime,_ = url.QueryUnescape(req.Form.Get("rb_DealTime"))
	notifyReq.RcBankCode = req.Form.Get("rc_BankCode")
	notifyReq.Hmac = req.Form.Get("hmac")
	//logging.Info(m.Get("ra_PayTime")+ m.Get("rb_DealTime"))
	if err != nil {
		return nil, err
	}
	return
}


func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}
