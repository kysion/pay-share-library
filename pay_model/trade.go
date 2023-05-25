package pay_model

type TradeRes struct {
	TradeNo string `json:"tradeNo" dc:"下单成功后返回的交易编号"`
	OrderId int64  `json:"orderId" dc:"订单id"`
}

type PayParamsRes struct { // 支付参数
	//// 支付宝
	//TradeNo *string `json:"tradeNo" dc:"Alipay下单成功后返回的交易编号"`

	// 微信
	AppId     *string `json:"appId" dc:"商户申请的公众号对应的appid，由微信支付生成，可在公众号后台查看。若下单时传了sub_appid，可为sub_appid的值" `
	TimeStamp *string `json:"timeStamp" dc:"时间戳，需要转换成秒(10位数字)"`
	NonceStr  *string `json:"nonceStr" dc:"随机字符串，不长于32位"`
	Package   *string `json:"package" dc:"JSAPI下单接口返回的prepay_id"`
	SignType  *string `json:"signType" dc:"签名类型，默认为RSA，仅支持RSA。"`
	PaySign   *string `json:"paySign" dc:"签名，使用字段appId、timeStamp、nonceStr、package计算得出的签名值"`
}
