package pay_enum

import (
    "github.com/kysion/pay-share-library/pay_model/pay_enum/internal/alipay_trade"
    "github.com/kysion/pay-share-library/pay_model/pay_enum/internal/order"
)

type (
    OrderAuditType order.AuditStateEnum
    OrderStateType order.StateTypeEnum

    // AlipayTradeStatus 阿里支付状态
    AlipayTradeStatus alipay_trade.TradeStatusEnum
)

var (

    // Order 订单
    Order = order.Order

    // AlipayTrade 阿里支付
    AlipayTrade = alipay_trade.AlipayTrade
)
