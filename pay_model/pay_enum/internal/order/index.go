package order

type order struct {
    StateType  stateType
    AuditState auditState

    TradeSourceType tradeSourceType
}

var Order = order{
    StateType:       StateType,
    AuditState:      AuditState,
    TradeSourceType: TradeSourceType,
}
