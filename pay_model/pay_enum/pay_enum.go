package pay_enum

import (
    "github.com/kysion/pay-share-library/pay_model/pay_enum/internal/order"
)

type (
    OrderAuditType order.AuditStateEnum
    OrderStateType order.StateTypeEnum
)

var (

    // Order 订单
    Order = order.Order
)

