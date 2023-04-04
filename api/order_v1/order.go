package order_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	model "github.com/kysion/pay-share-library/pay_model"
)

type GetOrderByIdReq struct {
	g.Meta `path:"/getOrderById" method:"post" summary:"根据ID查询订单|信息" tags:"订单"`
	Id     int64 `json:"id"              description:"id" v:"required#订单id不能为空"`
}

type GetOrderByPlatformOrderIdReq struct {
	g.Meta          `path:"/getOrderByPlatformOrderId" method:"post" summary:"根据第三平台交易id查询订单|信息" tags:"订单"`
	PlatformOrderId string `json:"id"              description:"id" v:"required#第三方平台交易id不能为空"`
}

type QueryOrderListReq struct {
	g.Meta `path:"/queryOrderList" method:"post" summary:"查询订单|列表" tags:"订单"`
	base_model.SearchParams
}

type QueryOrderByOneMonthReq struct {
	g.Meta `path:"/queryOrderByOneMonth" method:"post" summary:"查询最近一个月的订单|列表" tags:"订单"`
	base_model.SearchParams
}

type QueryOrderByTowMonthReq struct {
	g.Meta `path:"/queryOrderByTowMonth" method:"post" summary:"查询最近两个月的订单|列表" tags:"订单"`
	base_model.SearchParams
}

type AuditOrderRefundReq struct {
	g.Meta `path:"/auditOrderRefund" method:"post" summary:"审核订单退款|信息" tags:"订单"`
	model.AuditOrder
}

type UpdateOrderStateReq struct {
	g.Meta `path:"/updateOrderState" method:"post" summary:"修改订单状态|信息" tags:"订单"`
	Id     int64 `json:"id"              description:"id" v:"required#设备id不能为空"`
	State  int   `json:"state" dc:"订单状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭"`
}

type QueryOrderByProductNumberReq struct {
	g.Meta `path:"/queryOrderByProductNumber" method:"post" summary:"根据产品编号查询订单|列表" tags:"订单"`
	Number string `json:"number" dc:"产品编号就是设备编号" v:"required#产品编号不能为空"`
	base_model.SearchParams
}

type QueryOrderByUnionMainIdReq struct {
	g.Meta      `path:"/queryOrderByUnionMainId" method:"post" summary:"根据主体查询订单|列表" tags:"订单"`
	UnionMainId int64 `json:"union_main_id"              description:"主体ID，一般是app应用关联的主体" v:"required#主体id不能为空"`
	base_model.SearchParams
}

type GetOrderByConsumerIdReq struct {
	g.Meta `path:"/getOrderByConsumerId" method:"post" summary:"根据消费者查询订单|列表" tags:"订单"`
	Id     int64 `json:"id" dc:"消费者Id" v:"required#消费者id不能为空" `
	base_model.SearchParams
}
