package order

import "github.com/kysion/base-library/utility/enum"

// StateTypeEnum 订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
type StateTypeEnum enum.IEnumCode[int]

type stateType struct {
    Unpaid         StateTypeEnum
    During         StateTypeEnum
    HavePaid       StateTypeEnum
    OffPaid        StateTypeEnum
    DealClose      StateTypeEnum
    Refunding      StateTypeEnum
    Refunded       StateTypeEnum
    PaymentTimeOut StateTypeEnum
    Closed         StateTypeEnum
}

var StateType = stateType{
    Unpaid:         enum.New[StateTypeEnum](1, "待支付"),
    During:         enum.New[StateTypeEnum](2, "支付中"),
    HavePaid:       enum.New[StateTypeEnum](4, "已支付"),
    OffPaid:        enum.New[StateTypeEnum](8, "取消支付"),
    DealClose:      enum.New[StateTypeEnum](16, "交易完成"),
    Refunding:      enum.New[StateTypeEnum](32, "退款中"),
    Refunded:       enum.New[StateTypeEnum](64, "已退款"),
    PaymentTimeOut: enum.New[StateTypeEnum](128, "支付超时"),
    Closed:         enum.New[StateTypeEnum](256, "已关闭"),
}

func (e stateType) New(code int) StateTypeEnum {
    if (code & StateType.Unpaid.Code()) == StateType.Unpaid.Code() {
        return StateType.Unpaid
    }
    if (code & StateType.During.Code()) == StateType.During.Code() {
        return StateType.During
    }
    if (code & StateType.HavePaid.Code()) == StateType.HavePaid.Code() {
        return StateType.HavePaid
    }
    if (code & StateType.OffPaid.Code()) == StateType.OffPaid.Code() {
        return StateType.OffPaid
    }
    if (code & StateType.DealClose.Code()) == StateType.DealClose.Code() {
        return StateType.DealClose
    }
    if (code & StateType.Refunding.Code()) == StateType.Refunding.Code() {
        return StateType.Refunding
    }
    if (code & StateType.Refunded.Code()) == StateType.Refunded.Code() {
        return StateType.Refunded
    }
    if (code & StateType.PaymentTimeOut.Code()) == StateType.PaymentTimeOut.Code() {
        return StateType.PaymentTimeOut
    }
    if (code & StateType.Closed.Code()) == StateType.Closed.Code() {
        return StateType.Closed
    }

    panic("order.StateType.New: error")
}
