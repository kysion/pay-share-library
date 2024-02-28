// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package pay_service

import (
	"context"

	"github.com/kysion/base-library/base_model"
	model "github.com/kysion/pay-share-library/pay_model"
	enum "github.com/kysion/pay-share-library/pay_model/pay_enum"
	hook "github.com/kysion/pay-share-library/pay_model/pay_hook"
)

type (
	IOrder interface {
		// InstallHook 安装Hook的时候，如果状态类型为退款中，需要做响应的退款操作，谨防多模块订阅退款状态，产生重复退款
		InstallHook(actionType enum.OrderStateType, hookFunc hook.OrderHookFunc)
		// CreateOrder 创建订单
		CreateOrder(ctx context.Context, info *model.Order) (*model.OrderRes, error)
		// GetOrderById 根据ID查询订单
		GetOrderById(ctx context.Context, id int64) (*model.OrderRes, error)
		// GetOrderByPlatformOrderId 根据第三平台交易id查询订单  支付宝 == trade_no
		GetOrderByPlatformOrderId(ctx context.Context, id string) (*model.OrderRes, error)
		// QueryOrderList 查询订单列表
		QueryOrderList(ctx context.Context, info *base_model.SearchParams) (*model.OrderListRes, error)
		// QueryOrderByOneMonth 查询最近30天的订单
		QueryOrderByOneMonth(ctx context.Context, info *base_model.SearchParams) (*model.OrderListRes, error)
		// QueryOrderByTwoMonth 查询最近60天个月的订单
		QueryOrderByTwoMonth(ctx context.Context, info *base_model.SearchParams) (*model.OrderListRes, error)
		// AuditOrderRefund 审核订单退款   审核触发条件 --- 退款 或者取消支付
		AuditOrderRefund(ctx context.Context, info *model.AuditOrder) (bool, error)
		// UpdateOrderTradeSource 修改订单支付元数据
		UpdateOrderTradeSource(ctx context.Context, orderId int64, info *model.UpdateOrderTradeInfo) (bool, error)
		// UpdateOrderState 修改订单状态
		UpdateOrderState(ctx context.Context, id int64, state int) (bool, error)
		// GetOrderByProductNumber 根据产品编号查询订单
		GetOrderByProductNumber(ctx context.Context, number string, unionMainId int64) (*model.OrderRes, error)
		// GetLatestOrderByProductNumber  根据产品编号查询最新的一条订单
		GetLatestOrderByProductNumber(ctx context.Context, number string) (*model.OrderRes, error)
		// QueryOrderByProductNumber 根据产品编号查询订单|列表
		QueryOrderByProductNumber(ctx context.Context, number string, info *base_model.SearchParams, unionMainId int64, merchantId int64) (*model.OrderListRes, error)
		// GetOrderByUnionMainId 根据主体查询订单列表
		GetOrderByUnionMainId(ctx context.Context, id int64, info *base_model.SearchParams) (*model.OrderListRes, error)
		// GetOrderByConsumerId 根据消费者查询订单列表
		GetOrderByConsumerId(ctx context.Context, id int64, info *base_model.SearchParams) (*model.OrderListRes, error)
		// HasInUserOrder 根据设备ID判断是够存在正在使用的订单
		HasInUserOrder(ctx context.Context, deviceNumber string) (bool, error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
