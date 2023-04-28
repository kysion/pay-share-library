package pay_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type Order struct {
	Id               int64       `json:"id"              dc:"id"`
	PlatformOrderId  string      `json:"platformOrderId" dc:"第三方平台交易id，例如支付宝、微信..."`
	ConsumerId       int64       `json:"consumerId"      dc:"消费者id"`
	InOutType        int         `json:"inOutType"       dc:"收支类型：1收入，2支出" v:"required#收支类型不能为空"`
	Amount           int         `json:"amount"          dc:"交易金额，也就是实际成交金额" `
	CouponAmount     int         `json:"couponAmount"    dc:"优惠减免金额"`
	CouponConfig     string      `json:"couponConfig"    dc:"优惠配置详细数据"`
	OrderAmount      int         `json:"orderAmount"     dc:"订单金额，也就是优惠前的金额" v:"required#订单金额不能为空"`
	ProductName      string      `json:"productName"     dc:"产品名称，例如充电等" v:"required#产品名称不能为空"`
	TradeScene       string      `json:"tradeScene"      dc:"交易场景"`
	CreatedAt        *gtime.Time `json:"createdAt"       dc:"订单创建时间"`
	RefundAmount     int         `json:"refundAmount"    dc:"退款金额"`
	TradeAt          *gtime.Time `json:"tradeAt"         dc:"交易时间"`
	State            int         `json:"state"            dc:"订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭、512待分账、"`
	AuditState       int         `json:"auditState"      dc:"审核状态：0待审核、1已通过、-1不通过"`
	AuditReplyMsg    string      `json:"auditReplyMsg"  dc:"审核回复，审核状态为不通过时必须要有审核回复"`
	AuditAt          *gtime.Time `json:"auditAt"         dc:"审核时间"`
	TradeSourceType  int         `json:"tradeSourceType" dc:"交易源类型：1支付宝、2微信、4抖音、8银联"`
	TradeSource      string      `json:"tradeSource"     dc:"交易元数据"`
	ProductNumber    string      `json:"productNumber"   dc:"产品编号"  v:"required#产品编号不能为空"`
	UnionMainId      int64       `json:"unionMainId"      dc:"应用关联主体ID"`
	SubAccountScheme string      `json:"subAccountScheme" dc:"分账方案，默认空，空代表分账结束，无需分账"`
	AppId            string      `json:"appId"            dc:"应用appId"`
	UnionMainType    int         `json:"unionMainType"    dc:"应用关联主体类型, 也就是userType的类型"`
	MerchantId       int64       `json:"merchantId"       dc:"商家id"`
}

type CreateOrder struct {
	InOutType   int   `json:"inOutType"       dc:"收支类型：1收入，2支出"`
	Amount      int   `json:"amount"          dc:"交易金额，也就是实际成交金额"`
	OrderAmount int   `json:"orderAmount"     dc:"订单金额，也就是优惠前的金额"`
	State       int   `json:"state"           dc:"订单状态：1待支付、2支付中、4已支付、8支付失败、16交易完成、"`
	AccountId   int64 `json:"accountId"       dc:"财务账号id"`
}

type AuditOrder struct {
	Id    int64 `json:"id"                   dc:"id"`
	State int   `json:"state" dc:"订单状态：8取消支付，32退款"`
	//State          int    `json:"state"           dc:"订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭"`
	AuditState    int    `json:"auditState"      dc:"审核状态：0待审核、1已通过、-1不通过"`
	AuditReplyMsg string `json:"auditReplyMsg"  dc:"审核回复，审核状态为不通过时必须要有审核回复"`
}

type UpdateOrderTradeInfo struct {
	Id int64 `json:"id"              dc:"id"`

	PlatformOrderId string `json:"platformOrderId" dc:"第三方平台交易id，例如支付宝、微信..."`

	TradeSource string `json:"tradeSource"     dc:"交易元数据"`
}

type OrderRes Order

type OrderListRes base_model.CollectRes[Order]
