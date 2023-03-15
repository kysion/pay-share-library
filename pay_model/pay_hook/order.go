package pay_hook

import (
	"context"
	"github.com/kysion/pay-share-library/pay_model"
	"github.com/kysion/pay-share-library/pay_model/pay_enum"
)

// OrderHookFunc 上下文，状态，订单数据
type OrderHookFunc func(ctx context.Context, state pay_enum.OrderStateType, device pay_model.Order) bool

// OrderHookInfo Key是事件触发条件  Value是执行函数
type OrderHookInfo struct {
	Key   pay_enum.OrderStateType
	Value OrderHookFunc
}
