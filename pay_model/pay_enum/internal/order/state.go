package order

import "github.com/kysion/base-library/utility/enum"

// StateTypeEnum 订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
type StateTypeEnum enum.IEnumCode[int]

type stateType struct {
	WaitPayment     StateTypeEnum
	Paymenting      StateTypeEnum
	Paymented       StateTypeEnum
	CancelPayment   StateTypeEnum
	PaymentComplete StateTypeEnum
	Refunding       StateTypeEnum
	Refunded        StateTypeEnum
	PaymentTimeOut  StateTypeEnum
	ClosedPayment   StateTypeEnum
}

var StateType = stateType{
	WaitPayment:     enum.New[StateTypeEnum](1, "待支付"),
	Paymenting:      enum.New[StateTypeEnum](2, "支付中"),
	Paymented:       enum.New[StateTypeEnum](4, "已支付"),
	CancelPayment:   enum.New[StateTypeEnum](8, "取消支付"),
	PaymentComplete: enum.New[StateTypeEnum](16, "交易完成"),
	Refunding:       enum.New[StateTypeEnum](32, "退款中"),
	Refunded:        enum.New[StateTypeEnum](64, "已退款"),
	PaymentTimeOut:  enum.New[StateTypeEnum](128, "支付超时"),
	ClosedPayment:   enum.New[StateTypeEnum](256, "已关闭"),
}

func (e stateType) New(code int) StateTypeEnum {
	if (code & StateType.WaitPayment.Code()) == StateType.WaitPayment.Code() {
		return StateType.WaitPayment
	}
	if (code & StateType.Paymenting.Code()) == StateType.Paymenting.Code() {
		return StateType.Paymenting
	}
	if (code & StateType.Paymented.Code()) == StateType.Paymented.Code() {
		return StateType.Paymented
	}
	if (code & StateType.CancelPayment.Code()) == StateType.CancelPayment.Code() {
		return StateType.CancelPayment
	}
	if (code & StateType.PaymentComplete.Code()) == StateType.PaymentComplete.Code() {
		return StateType.PaymentComplete
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
	if (code & StateType.ClosedPayment.Code()) == StateType.ClosedPayment.Code() {
		return StateType.ClosedPayment
	}

	panic("order.StateType.New: error")
}
