//==================================
//  * Name：Michael
//  * Tel：15014401199
//  * DateTime：2019/09/29 14:58
//==================================
package joinpay

type UniPayApiResponse struct {
	R0Version 			string `json:"r0_Version"`
	R1MerchantNo  		string `json:"r1_MerchantNo"`
	R2OrderNo      		string `json:"r2_OrderNo"`
	R3Amount      		string `json:"r3_Amount"`
	R4Cur 				string `json:"r4_Cur"`
	R5Mp   				string `json:"r5_Mp"`
	R6FrpCode       	string `json:"r6_FrpCode"`
	R7TrxNo 			string `json:"r7_TrxNo"`
	R8MerchantBankCode  string `json:"r8_MerchantBankCode"`
	R9SubMerchantNo 	string `json:"r9_SubMerchantNo"`
	RaCode  			int64  `json:"ra_Code"`
	RbCodeMsg   		string `json:"rb_CodeMsg"`
	RcResult    		string `json:"rc_Result"`
	RdPic    			string `json:"rd_Pic"`
	Hmac    			string `json:"hmac"`
}

type JoinNotifyRequest struct {
	R1MerchantNo        string `json:"r1_MerchantNo"`
	R2OrderNo          	string `json:"r2_OrderNo"`
	R3Amount         	string `json:"r3_Amount"`
	R4Cur            	string `json:"r4_Cur"`
	R5Mp         		string `json:"r5_Mp"`
	R6Status            string `json:"r6_Status"`
	R7TrxNo             string `json:"r7_TrxNo"`
	R8BankOrderNo       string `json:"r8_BankOrderNo"`
	R9BankTrxNo         string `json:"r9_BankTrxNo"`
	RaPayTime           string `json:"ra_PayTime"`
	RbDealTime          string `json:"rb_DealTime"`
	RcBankCode          string `json:"rc_BankCode"`
	Hmac        		string `json:"hmac"`
}


type RefundResponse struct {
	R1MerchantNo		string `json:"r1_MerchantNo"`
	R2OrderNo			string `json:"r2_OrderNo"`
	R3RefundOrderNo		string `json:"r3_RefundOrderNo"`
	R4RefundAmount		string `json:"r4_RefundAmount"`
	RaStatus			string `json:"ra_Status"`
	RbCode				string `json:"rb_Code"`
	RcCodeMsg			string `json:"rc_CodeMsg"`
	Hmac				string `json:"hmac"`
}


type RefundNotifyResponse struct {
	R1MerchantNo		string `json:"r1_MerchantNo"`
	R2OrderNo			string `json:"r2_OrderNo"`
	R3RefundOrderNo		string `json:"r3_RefundOrderNo"`
	R4RefundAmountStr	string `json:"r4_RefundAmount_str"`
	RaStatus			string `json:"ra_Status"`
	Hmac				string `json:"hmac"`
}