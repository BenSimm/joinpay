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

