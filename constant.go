package joinpay

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"

	//URL
	wx_base_url        = "https://www.joinpay.com/"

	//正式
	joinURL_UniPayApi  = wx_base_url + "trade/uniPayApi.action"          //请求交易入口

	//支付类型
	TradeType_ALi_Native   		= "ALIPAY_NATIVE"
	TradeType_Ali_Card  		= "ALIPAY_CARD"
	TradeType_ALi_App    		= "ALIPAY_APP"
	TradeType_Ali_H5     		= "ALIPAY_H5"
	TradeType_Ali_Fwc 			= "ALIPAY_FWC"
	TradeType_Ali_Syt 			= "ALIPAY_SYT"
	TradeType_Wx_Native 		= "WEIXIN_NATIVE"
	TradeType_Wx_Card 			= "WEIXIN_CARD"
	TradeType_Wx_App 			= "WEIXIN_APP"
	TradeType_Wx_H5 			= "WEIXIN_H5"
	TradeType_Wx_Gzh 			= "WEIXIN_GZH"
	TradeType_Wx_Xcx 			= "WEIXIN_XCX"
	TradeType_Jd_Native 		= "JD_NATIVE"
	TradeType_Jd_Card 			= "JD_CARD"
	TradeType_Jd_App 			= "JD_APP"
	TradeType_Jd_H5 			= "JD_H5"
	TradeType_Qq_Native 		= "QQ_NATIVE"
	TradeType_Qq_Card 			= "QQ_CARD"
	TradeType_Qq_App 			= "QQ_APP"
	TradeType_Qq_H5 			= "QQ_H5"
	TradeType_Qq_Gzh 			= "QQ_GZH"
	TradeType_Unionpay_Native 	= "UNIONPAY_NATIVE"
	TradeType_Unionpay_Card 	= "UNIONPAY_CARD"
	TradeType_Unionpay_App 		= "UNIONPAY_APP"
	TradeType_Unionpay_H5 		= "UNIONPAY_H5"
	TradeType_Baidu_Native 		= "BAIDU_NATIVE"
	TradeType_Suning_Native 	= "SUNING_NATIVE"

	//签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
)
