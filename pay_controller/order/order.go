package order

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/pay-share-library/api/order_v1"
	model "github.com/kysion/pay-share-library/pay_model"
	service "github.com/kysion/pay-share-library/pay_service"
)

var Order = cOrder{}

type cOrder struct{}

func (c *cOrder) GetOrderById(ctx context.Context, req *order_v1.GetOrderByIdReq) (*model.OrderRes, error) {
	ret, err := service.Order().GetOrderById(ctx, req.Id)

	return ret, err
}

func (c *cOrder) GetOrderByPlatformOrderId(ctx context.Context, req *order_v1.GetOrderByPlatformOrderIdReq) (*model.OrderRes, error) {
	ret, err := service.Order().GetOrderByPlatformOrderId(ctx, req.PlatformOrderId)

	return ret, err
}

func (c *cOrder) QueryOrderList(ctx context.Context, req *order_v1.QueryOrderListReq) (*model.OrderListRes, error) {
	ret, err := service.Order().QueryOrderList(ctx, &req.SearchParams)

	return ret, err
}

func (c *cOrder) QueryOrderByOneMonth(ctx context.Context, req *order_v1.QueryOrderByOneMonthReq) (*model.OrderListRes, error) {
	ret, err := service.Order().QueryOrderByOneMonth(ctx, &req.SearchParams)

	return ret, err
}

func (c *cOrder) QueryOrderByTwoMonth(ctx context.Context, req *order_v1.QueryOrderByTowMonthReq) (*model.OrderListRes, error) {
	ret, err := service.Order().QueryOrderByTwoMonth(ctx, &req.SearchParams)

	return ret, err
}

func (c *cOrder) AuditOrderRefund(ctx context.Context, req *order_v1.AuditOrderRefundReq) (api_v1.BoolRes, error) {
	ret, err := service.Order().AuditOrderRefund(ctx, &req.AuditOrder)

	return ret == true, err
}

func (c *cOrder) UpdateOrderState(ctx context.Context, req *order_v1.UpdateOrderStateReq) (api_v1.BoolRes, error) {
	ret, err := service.Order().UpdateOrderState(ctx, req.Id, req.State)

	return ret == true, err
}

func (c *cOrder) QueryOrderByProductNumber(ctx context.Context, req *order_v1.QueryOrderByProductNumberReq) (*model.OrderListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := service.Order().QueryOrderByProductNumber(ctx, req.Number, &req.SearchParams, user.UnionMainId, 0)

	return ret, err
}

func (c *cOrder) GetOrderByUnionMainId(ctx context.Context, req *order_v1.QueryOrderByUnionMainIdReq) (*model.OrderListRes, error) {
	ret, err := service.Order().GetOrderByUnionMainId(ctx, req.UnionMainId, &req.SearchParams)

	return ret, err
}

func (c *cOrder) GetOrderByConsumerId(ctx context.Context, req *order_v1.GetOrderByConsumerIdReq) (*model.OrderListRes, error) {
	ret, err := service.Order().GetOrderByConsumerId(ctx, req.Id, &req.SearchParams)

	return ret, err
}
