package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/backend"
	"my_shop/internal/consts"

	"my_shop/internal/model"
	"my_shop/internal/service"
)

// 承上启下  mvc
// Admin 管理员管理
var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{AdminId: out.AdminId}, nil
}

func (a *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

func (a *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return
}

func (a *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	return &backend.AdminGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

//for jwt
//func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
//	return &backend.AdminGetInfoRes{
//		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
//		IdentityKey: service.Auth().IdentityKey,
//		Payload:     service.Auth().GetPayload(ctx),
//	}, nil
//}

// Info for gtoken
func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
	return &backend.AdminGetInfoRes{
		Id:      gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		RoleIds: gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxAdminIsAdmin)),
	}, nil
}

//func (a *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
//	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &backend.AdminGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}

// 前台的取值方法
//func (a *cAdmin) ListFrontend(ctx context.Context, req *frontend.AdminGetListCommonReq) (res *frontend.AdminGetListCommonRes, err error) {
//	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &frontend.AdminGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}
