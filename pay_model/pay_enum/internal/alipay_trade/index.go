package alipay_trade

type alipayTrade struct {
    TradeStatus tradeStatus

    OperationType operationType
}

var AlipayTrade = alipayTrade{
    TradeStatus: TradeStatus,

    OperationType: OperationType,
}
