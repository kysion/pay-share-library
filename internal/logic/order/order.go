package order

import (
    "context"
    "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
    "github.com/SupenBysz/gf-admin-community/sys_service"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gtime"
    "github.com/gogf/gf/v2/util/gconv"
    "github.com/kysion/base-library/base_hook"
    "github.com/kysion/base-library/base_model"
    "github.com/kysion/base-library/utility/daoctl"
    "github.com/kysion/base-library/utility/kconv"
    model "github.com/kysion/pay-share-library/pay_model"
    dao "github.com/kysion/pay-share-library/pay_model/pay_dao"
    do "github.com/kysion/pay-share-library/pay_model/pay_do"

    enum "github.com/kysion/pay-share-library/pay_model/pay_enum"
    hook "github.com/kysion/pay-share-library/pay_model/pay_hook"
    "github.com/kysion/pay-share-library/pay_service"
    "github.com/yitter/idgenerator-go/idgen"

    "time"
)

// 域名 路径 具体的接口地址

type sOrder struct {
    Duration time.Duration
    base_hook.BaseHook[enum.OrderStateType, hook.OrderHookFunc]
}

func init() {
    pay_service.RegisterOrder(New())
}

func New() *sOrder {
    return &sOrder{}
}

// InstallHook 安装Hook的时候，如果状态类型为退款中，需要做响应的退款操作，谨防多模块订阅退款状态，产生重复退款
func (s *sOrder) InstallHook(actionType enum.OrderStateType, hookFunc hook.OrderHookFunc) {
    s.BaseHook.InstallHook(actionType, hookFunc)
}

// CreateOrder 创建订单
func (s *sOrder) CreateOrder(ctx context.Context, info *model.Order) (*model.OrderRes, error) {
    if info.Id == 0 {
        info.Id = idgen.NextId()
    }

    data := do.Order{}

    gconv.Struct(info, &data)
    // 订单状态默认是待支付
    if info.State == 0 {
        data.State = 1
    }
    if info.CouponConfig == "" {
        data.CouponConfig = nil
    }
    if info.TradeScene == "" {
        data.TradeScene = nil
    }
    if info.TradeSource == "" {
        data.TradeSource = nil
    }
	
    data.TradeAt = gtime.Now()

    affected, err := daoctl.InsertWithError(dao.Order.Ctx(ctx), data)

    if affected == 0 || err != nil {
        return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "订单添加失败", dao.Order.Table())
    }

    return s.GetOrderById(ctx, gconv.Int64(data.Id))
}

// GetOrderById 根据ID查询订单
func (s *sOrder) GetOrderById(ctx context.Context, id int64) (*model.OrderRes, error) {
    return daoctl.GetByIdWithError[model.OrderRes](dao.Order.Ctx(ctx), id)
}

// GetOrderByPlatformOrderId 根据第三平台交易id查询订单  支付宝 == trade_no
func (s *sOrder) GetOrderByPlatformOrderId(ctx context.Context, id string) (*model.OrderRes, error) {
    user := sys_service.SysSession().Get(ctx).JwtClaimsUser

    daoWhere := dao.Order.Ctx(ctx).Where(do.Order{PlatformOrderId: id})

    //sys_enum.User.Type.SuperAdmin
    if (user.Type & sys_enum.User.Type.Admin.Code()) != sys_enum.User.Type.Admin.Code() {
        daoWhere = daoWhere.Where(do.Order{UnionMainId: user.UnionMainId}).WhereOr(do.Order{MerchantId: user.UnionMainId})   // 商户和应用的所属商家需要去看订单
    }

    data, err := daoctl.ScanWithError[model.OrderRes](daoWhere)

    if err != nil {
        return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据第三方平台交易id查询订单失败，请检查", dao.Order.Table())
    }

    return data, nil
}

// QueryOrderList 查询订单列表
func (s *sOrder) QueryOrderList(ctx context.Context, info *base_model.SearchParams) (*model.OrderListRes, error) {
    //user := sys_service.SysSession().Get(ctx).JwtClaimsUser

    daoWhere := dao.Order.Ctx(ctx)

    //if (user.Type & sys_enum.User.Type.Admin.Code()) != sys_enum.User.Type.Admin.Code() {
    //	daoWhere = daoWhere.Where(do.Order{UnionMainId: user.UnionMainId})
    //}

    result, err := daoctl.Query[model.Order](daoWhere, info, false)

    return (*model.OrderListRes)(result), err
}

// QueryOrderByOneMonth 查询最近30天的订单
func (s *sOrder) QueryOrderByOneMonth(ctx context.Context) (*model.OrderListRes, error) {
    user := sys_service.SysSession().Get(ctx).JwtClaimsUser

    now := gtime.Now()

    time := now.Add(0 - time.Hour*24*30)

    daoWhere := dao.Order.Ctx(ctx).WhereGT(dao.Order.Columns().CreatedAt, time)

    if (user.Type & sys_enum.User.Type.Admin.Code()) != sys_enum.User.Type.Admin.Code() {
        daoWhere = daoWhere.Where(do.Order{UnionMainId: user.UnionMainId}).WhereOr(do.Order{MerchantId: user.UnionMainId})  // 商户和应用的所属商家需要去看订单
    }

    data, err := daoctl.Query[model.Order](daoWhere, nil, false)

    if err != nil {
        return &model.OrderListRes{}, err
    }

    return (*model.OrderListRes)(data), nil
}

// QueryOrderByTwoMonth 查询最近60天个月的订单
func (s *sOrder) QueryOrderByTwoMonth(ctx context.Context) (*model.OrderListRes, error) {
    user := sys_service.SysSession().Get(ctx).JwtClaimsUser

    now := gtime.Now()

    time := now.Add(0 - time.Hour*24*60)

    daoWhere := dao.Order.Ctx(ctx).WhereGT(dao.Order.Columns().CreatedAt, time)

    if (user.Type & sys_enum.User.Type.Admin.Code()) != sys_enum.User.Type.Admin.Code() {
        daoWhere = daoWhere.Where(do.Order{UnionMainId: user.UnionMainId}).WhereOr(do.Order{MerchantId: user.UnionMainId})   // 商户和应用的所属商家需要去看订单
    }

    data, err := daoctl.Query[model.Order](daoWhere, nil, false)

    if err != nil {
        return &model.OrderListRes{}, err
    }

    return (*model.OrderListRes)(data), nil
}

// AuditOrderRefund 审核订单退款   审核触发条件 --- 退款 或者取消支付
func (s *sOrder) AuditOrderRefund(ctx context.Context, info *model.AuditOrder) (bool, error) {
    orderInfo, err := s.GetOrderById(ctx, info.Id)
    if orderInfo == nil {
        return false, sys_service.SysLogs().ErrorSimple(ctx, err, "审核订单不存在", dao.Order.Table())
    }
    // 代表已审过的
    if orderInfo.AuditState > enum.Order.AuditState.WaitAudit.Code() {
        return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为错误", dao.Order.Table())
    }

    // 只能审核待审核的订单
    if (orderInfo.AuditState & enum.Order.AuditState.WaitAudit.Code()) != enum.Order.AuditState.WaitAudit.Code() {
        return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "只能审核待审核的订单", dao.Order.Table())
    }

    // 审核不通过必须要有原因说明
    if info.AuditState == enum.Order.AuditState.Reject.Code() && info.AuditReplyMsg == "" {
        return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过必须要有原因说明", dao.Order.Table())
    }

    // 退款
    order := do.Order{
        AuditReplyMsg: info.AuditReplyMsg,
        AuditAt:       gtime.Now(),
        AuditState:    info.AuditState,
        State:         enum.Order.StateType.Refunded.Code(), // 已退款
    }

    // 审核不通过后订单状态还是退款中
    if info.AuditState == enum.Order.AuditState.Reject.Code() {
        order.State = enum.Order.StateType.Refunded.Code()
    }

    // 退款的话需要执行退款流程 需要补充在具体的业务Hook执行退款逻辑
    s.Iterator(func(key enum.OrderStateType, value hook.OrderHookFunc) {

        // 如果注入的Hook的key是退款中，那么执行该注入Hook的Func
        if (key.Code() & enum.Order.StateType.Refunding.Code()) == enum.Order.StateType.Refunding.Code() {
            g.Try(ctx, func(ctx context.Context) {
                // 调用HookFunc
                value(ctx, enum.Order.StateType.Refunding, model.Order{Id: info.Id})

                // 修改审核状态,该行为在业务层进行处理

            })
        }
    })

    if err != nil {
        return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "订单审核失败", dao.Order.Table())
    }

    return true, nil
}

// UpdateOrderTradeSource 修改订单支付元数据
func (s *sOrder) UpdateOrderTradeSource(ctx context.Context, info *model.UpdateOrderTradeInfo) (bool, error) {
    // 先判断是否存在
    videoInfo, err := s.GetOrderById(ctx, info.Id)
    if err != nil || videoInfo == nil {
        return false, sys_service.SysLogs().ErrorSimple(ctx, err, "订单不存在", dao.Order.Table())
    }

    data := kconv.Struct(info, &do.Order{})

    affected, err := daoctl.UpdateWithError(dao.Order.Ctx(ctx).Data(data).OmitEmptyData().Where(do.Order{Id: info.Id}))

    if err != nil {
        return false, sys_service.SysLogs().ErrorSimple(ctx, err, "订单信息更新失败", dao.Order.Table())
    }

    return affected > 0, nil
}

// UpdateOrderState 修改订单状态
func (s *sOrder) UpdateOrderState(ctx context.Context, id int64, state int) (bool, error) {
    // 通过Hook解决不同的回调类型, 例如添加报损状态添加一条反馈记录
    _, err := s.GetOrderById(ctx, id)
    if err != nil {
        return false, err
    }

    info := do.Order{
        State: state,
    }

    // 如果是已支付，支付时间修改为现在
    if state == enum.Order.StateType.Paymented.Code() {
        info.TradeAt = gtime.Now()
    }

    // 通知订阅者订单状态发生了变化

    // 循环hookArr ，监听是否有人注册了Hook, 例如是一些特殊状态的时候需要做一些处理
    s.Iterator(func(key enum.OrderStateType, value hook.OrderHookFunc) {
        g.Try(ctx, func(ctx context.Context) {
            value(ctx, enum.Order.StateType.New(state), model.Order{Id: id})
        })
    })

    affected, err := daoctl.UpdateWithError(dao.Order.Ctx(ctx).Data(info).OmitEmptyData().Where(do.Order{Id: id}))

    if err != nil {
        return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设备状态修改失败", dao.Order.Table())
    }

    return affected > 0, err

}

// GetOrderByProductNumber 根据产品编号查询订单
func (s *sOrder) GetOrderByProductNumber(ctx context.Context, number string) (*model.OrderListRes, error) {
    user := sys_service.SysSession().Get(ctx).JwtClaimsUser

    daoWhere := dao.Order.Ctx(ctx).Where(dao.Order.Columns().ProductNumber, number)

    if (user.Type & sys_enum.User.Type.Admin.Code()) != sys_enum.User.Type.Admin.Code() {
        daoWhere = daoWhere.Where(do.Order{UnionMainId: user.UnionMainId}).WhereOr(do.Order{MerchantId: user.UnionMainId})  // 商户和应用的所属商家需要去看订单
    }

    data, err := daoctl.Query[model.Order](daoWhere, nil, false)

    if err != nil {
        return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据产品编号查询订单失败，请检查", dao.Order.Table())
    }

    return (*model.OrderListRes)(data), nil
}

// GetOrderByUnionMainId 根据主体查询订单列表
func (s *sOrder) GetOrderByUnionMainId(ctx context.Context, id int64) (*model.OrderListRes, error) {
    data, err := s.getOrderByAnyId(ctx, gconv.String(id), dao.Order.Columns().UnionMainId)

    if err != nil {
        return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据主体查询订单失败，请检查", dao.Order.Table())
    }

    return data, nil
}

// GetOrderByConsumerId 根据消费者查询订单列表
func (s *sOrder) GetOrderByConsumerId(ctx context.Context, id string) (*model.OrderListRes, error) {
    data, err := s.getOrderByAnyId(ctx, id, dao.Order.Columns().ConsumerId)

    if err != nil {
        return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据消费者查询订单失败，请检查", dao.Order.Table())
    }

    return data, nil
}

// 查询订单
func (s *sOrder) getOrderByAnyId(ctx context.Context, id string, columnName string) (*model.OrderListRes, error) {
    data, err := daoctl.Query[model.Order](dao.Order.Ctx(ctx).Where(columnName, id), nil, false)

    if err != nil {
        return nil, err
    }

    return (*model.OrderListRes)(data), nil
}
