package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/backend"
	"my_shop/internal/consts"
	"my_shop/internal/controller"
	"my_shop/internal/dao"
	"my_shop/internal/model/entity"
	"my_shop/internal/service"
	"my_shop/utility"
	"my_shop/utility/response"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				ServerName:       "myshop",
				CacheMode:        2, //redis
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/user/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true,
			}
			// todo 抽取方法
			err = gfAdminToken.Start()
			if err != nil {
				return err
			}
			// 认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(
				//	ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// gtoken中间件绑定
				//err := gfAdminToken.Middleware(ctx, group)
				//if err != nil {
				//	panic(err)
				//}
				group.Bind(
					controller.Hello,
					controller.Rotation,     //轮播图
					controller.Position,     //手工位
					controller.Admin.Create, //管理员
					controller.Admin.Update, //管理员
					controller.Admin.Delete, //管理员
					controller.Admin.List,   //管理员
					controller.Login,        //登录
					controller.Data,         //数据大屏
					controller.Role,         //角色
					controller.Permission,   // 权限
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) //for jwt
					// gtoken中间件绑定
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,   // 从0到1实现文件入库
						controller.Upload, // 实现可跨项目使用的文件上云工具类
					)
				})
			})
			s.Run()
			return nil
		},
	}
)

// todo 迁移到合适的位置
func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 调试打印
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// todo 迁移到合适的位置
// loginAfterFunc 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		g.Dump("admin:", adminId)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    10 * 24 * 60 * 60, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.Middleware.Next()
}
