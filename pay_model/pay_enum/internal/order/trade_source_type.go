package order

import "github.com/kysion/base-library/utility/enum"

type TradeSourceTypeEnum enum.IEnumCode[int]

type tradeSourceType struct {
	Alipay   TradeSourceTypeEnum
    Weixin   TradeSourceTypeEnum
    Douyin   TradeSourceTypeEnum
    Unionpay TradeSourceTypeEnum
}

var TradeSourceType = tradeSourceType{
	Alipay:   enum.New[TradeSourceTypeEnum](1, "支付宝"),
	Weixin :   enum.New[TradeSourceTypeEnum](2, "微信"),
    Douyin:   enum.New[TradeSourceTypeEnum](4, "抖音"),
    Unionpay: enum.New[TradeSourceTypeEnum](8, "银联"),
}

func (e tradeSourceType) New(code int, description string) TradeSourceTypeEnum {
    if (code&TradeSourceType.Alipay.Code()) == TradeSourceType.Alipay.Code()||
		(code&TradeSourceType.Weixin.Code()) == TradeSourceType.Weixin.Code() ||
        (code&TradeSourceType.Douyin.Code()) == TradeSourceType.Douyin.Code() ||
        (code&TradeSourceType.Unionpay.Code()) == TradeSourceType.Unionpay.Code() {
        return enum.New[TradeSourceTypeEnum](code, description)

    } else {
        panic("TradeSourceType.AuditState.New: error")
    }
}
