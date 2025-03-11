package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"my_shop/internal/consts"
	"my_shop/internal/controller"
	"my_shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动管理后台GToken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			// 认证接口
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(
				//	ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					controller.Hello,
					controller.Rotation.List,   //轮播图
					controller.Rotation.Delete, //轮播图
					controller.Rotation.Create, //轮播图
					controller.Rotation.Update, //轮播图
					controller.Position,        //手工位
					controller.Admin.Create,    //管理员
					controller.Admin.Update,    //管理员
					controller.Admin.Delete,    //管理员
					controller.Admin.List,      //管理员
					controller.Login,           //登录
					controller.Data,            //数据大屏
					controller.Role,            //角色
					controller.Permission,      // 权限
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) //for jwt
					// gtoken中间件绑定
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,         // 从0到1实现文件入库
						controller.Upload,       // 实现可跨项目使用的文件上云工具类
						controller.Category,     //商品分类管理
						controller.Coupon,       //优惠券管理
						controller.UserCoupon,   //用户优惠卷管理
						controller.Goods,        //商品管理
						controller.GoodsOptions, //商品规格管理
						controller.Article,      //文章管理
					)
				})
			})
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					controller.Rotation.ListFrontend, // 前台轮播图
				)
			})
			s.Run()
			return nil
		},
	}
)
