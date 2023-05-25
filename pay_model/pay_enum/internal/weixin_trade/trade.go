package weixin_trade

import "github.com/kysion/base-library/utility/enum"

type TradeStatusEnum enum.IEnumCode[string]

// 交易状态类型
// SUCCESS：支付成功
// REFUND：转入退款
// NOTPAY：未支付
// CLOSED：已关闭
// REVOKED：已撤销（仅付款码支付会返回）
// USERPAYING：用户支付中（仅付款码支付会返回）
// PAYERROR：支付失败（仅付款码支付会返回）

type tradeStatus struct {
	SUCCESS    TradeStatusEnum
	REFUND     TradeStatusEnum
	NOTPAY     TradeStatusEnum
	CLOSED     TradeStatusEnum
	REVOKED    TradeStatusEnum
	USERPAYING TradeStatusEnum
	PAYERROR   TradeStatusEnum
}

var TradeStatus = tradeStatus{
	SUCCESS:    enum.New[TradeStatusEnum]("SUCCESS", "支付成功"),
	REFUND:     enum.New[TradeStatusEnum]("REFUND", "转入退款"),
	NOTPAY:     enum.New[TradeStatusEnum]("NOTPAY", "未支付"),
	CLOSED:     enum.New[TradeStatusEnum]("CLOSED", "已关闭"),
	REVOKED:    enum.New[TradeStatusEnum]("REVOKED", "已撤销（仅付款码支付会返回）"),
	USERPAYING: enum.New[TradeStatusEnum]("USERPAYING", "用户支付中（仅付款码支付会返回）"),
	PAYERROR:   enum.New[TradeStatusEnum]("PAYERROR", "支付失败（仅付款码支付会返回）"),
}
