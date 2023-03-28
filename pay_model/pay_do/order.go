// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table kmk_order for DAO operations like Where/Data.
type Order struct {
	g.Meta           `orm:"table:kmk_order, do:true"`
	Id               interface{} // id
	PlatformOrderId  interface{} // 第三方平台交易id，例如支付宝、微信...
	ConsumerId       interface{} // 消费者id
	InOutType        interface{} // 收支类型：1收入，2支出
	Amount           interface{} // 交易金额，也就是实际成交金额
	CouponAmount     interface{} // 优惠减免金额
	CouponConfig     interface{} // 优惠配置详细数据
	OrderAmount      interface{} // 订单金额，也就是优惠前的金额
	ProductName      interface{} // 产品名称，例如充电等
	TradeScene       interface{} // 交易场景
	CreatedAt        *gtime.Time // 订单创建时间
	RefundAmount     interface{} // 退款金额
	TradeAt          *gtime.Time // 交易时间
	State            interface{} // 订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
	AuditState       interface{} // 审核状态：0待审核、1已通过、-1不通过
	AuditReplyMsg    interface{} // 审核回复，审核状态为不通过时必须要有审核回复
	AuditAt          *gtime.Time // 审核时间
	TradeSourceType  interface{} // 交易源类型：1支付宝、2微信、4抖音、8银联
	TradeSource      interface{} // 交易元数据
	ProductNumber    interface{} // 产品编号
	UnionMainId      interface{} // 应用关联主体ID
	SubAccountScheme interface{} // 分账方案，默认空，空代表分账结束，无需分账
	AppId            interface{} // 应用appId
	UnionMainType    interface{} // 应用关联主体类型, 也就是userType的类型
	MerchantId       interface{} // 商家id
}
