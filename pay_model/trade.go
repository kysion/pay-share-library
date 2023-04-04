package pay_model

type TradeRes struct {
	TradeNo string `json:"tradeNo" dc:"支付宝的交易编号"`
	OrderId int64  `json:"orderId" dc:"订单id"`
}
