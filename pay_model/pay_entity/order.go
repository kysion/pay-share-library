// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package share_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id              int64       `json:"id"              description:"id"`
	PlatformOrderId string      `json:"platformOrderId" description:"第三方平台交易id，例如支付宝、微信..."`
	ConsumerId      int64       `json:"consumerId"      description:"消费者id"`
	InOutType       int         `json:"inOutType"       description:"收支类型：1收入，2支出"`
	Amount          int         `json:"amount"          description:"交易金额，也就是实际成交金额"`
	CouponAmount    int         `json:"couponAmount"    description:"优惠减免金额"`
	CouponConfig    string      `json:"couponConfig"    description:"优惠配置详细数据"`
	OrderAmount     int         `json:"orderAmount"     description:"订单金额，也就是优惠前的金额"`
	BeforeBalance   int         `json:"beforeBalance"   description:"交易前余额"`
	AfterBalance    int         `json:"afterBalance"    description:"交易后余额"`
	ProductName     string      `json:"productName"     description:"产品名称，例如充电等"`
	TradeScene      string      `json:"tradeScene"      description:"交易场景"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"订单创建时间"`
	RefundAmount    int         `json:"refundAmount"    description:"退款金额"`
	TradeAt         *gtime.Time `json:"tradeAt"         description:"交易时间"`
	State           int         `json:"state"           description:"订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭"`
	AuditState      int         `json:"auditState"      description:"审核状态：0待审核、1已通过、-1不通过"`
	AuditReplayMsg  string      `json:"auditReplayMsg"  description:"审核回复，审核状态为不通过时必须要有审核回复"`
	AuditAt         *gtime.Time `json:"auditAt"         description:"审核时间"`
	TradeSourceType int         `json:"tradeSourceType" description:"交易源类型：1支付宝、2微信、4抖音、8银联"`
	TradeSource     string      `json:"tradeSource"     description:"交易元数据"`
	ProductNumber   string      `json:"productNumber"   description:"产品编号"`
	UnionMainId     int64       `json:"unionMainId"     description:"关联主体ID"`
}
