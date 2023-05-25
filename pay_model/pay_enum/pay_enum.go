package pay_enum

import (
	"github.com/kysion/pay-share-library/pay_model/pay_enum/internal/alipay_trade"
	"github.com/kysion/pay-share-library/pay_model/pay_enum/internal/order"
	"github.com/kysion/pay-share-library/pay_model/pay_enum/internal/weixin_trade"
)

type (
	OrderAuditType       order.AuditStateEnum
	OrderStateType       order.StateTypeEnum
	OrderTradeSourceType order.TradeSourceTypeEnum

	// AlipayTradeStatus 阿里支付状态
	AlipayTradeStatus alipay_trade.TradeStatusEnum

	// WeiXinTradeStatus 微信支付状态
	WeiXinTradeStatus weixin_trade.TradeStatusEnum
)

var (

	// Order 订单
	Order = order.Order

	// AlipayTrade 阿里支付
	AlipayTrade = alipay_trade.AlipayTrade

	// WeiXinTrade 微信支付
	WeiXinTrade = weixin_trade.WeiXinTrade
)
