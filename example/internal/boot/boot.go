package boot

import (
    "context"
    "github.com/SupenBysz/gf-admin-community/sys_controller"
    "github.com/SupenBysz/gf-admin-community/sys_service"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/os/gcmd"
    _ "github.com/kysion/pay-share-library/example/internal/boot/internal"
    "github.com/kysion/pay-share-library/pay_controller/order"
)

var (
    Main = gcmd.Command{
        Name:  "main",
        Usage: "main",
        Brief: "start http server",
        Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
            s := g.Server()

            s.Group("/alipay", func(group *ghttp.RouterGroup) {
                // 注册中间件
                group.Middleware(
                    sys_service.Middleware().CTX,
                    sys_service.Middleware().ResponseHandler,
                )

                // 不需要鉴权，但是需要登录的路由
                group.Group("/", func(group *ghttp.RouterGroup) {
                    // 注册中间件
                    group.Middleware(
                        sys_service.Middleware().Auth,
                    )
                    // 文件上传
                    group.Group("/common/sys_file", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysFile) })
                })

                // 匿名路由绑定
                group.Group("/", func(group *ghttp.RouterGroup) {
                    // 鉴权：登录，注册，找回密码等
                    group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
                    // 图型验证码、短信验证码、地区
                    group.Group("/common", func(group *ghttp.RouterGroup) {
                        group.Bind(
                            // 图型验证码
                            sys_controller.Captcha,
                            // 短信验证码
                            sys_controller.SysSms,
                            // 地区
                            sys_controller.SysArea,
                        )
                    })
                })

                // 订单模块
                group.Group("/order", func(group *ghttp.RouterGroup) { group.Bind(order.Order) })

            })

            s.Run()
            return nil
        },
    }
)
