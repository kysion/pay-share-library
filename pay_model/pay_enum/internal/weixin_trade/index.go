package weixin_trade

type weixinTrade struct {
	TradeStatus tradeStatus

	OperationType operationType
}

var WeiXinTrade = weixinTrade{
	TradeStatus: TradeStatus,

	OperationType: OperationType,
}
