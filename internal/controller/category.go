package controller

import (
	"context"
	"my_shop/api/backend"

	"my_shop/internal/model"
	"my_shop/internal/service"
)

// 承上启下  mvc
// Category 商品分类

var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryRes{CategoryId: out.CategoryId}, nil
}

func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: req.Id,
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
		},
	})
	return
}

func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	return &backend.CategoryGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

//func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
//	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &backend.CategoryGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}

// 前台的取值方法
//func (a *cCategory) ListFrontend(ctx context.Context, req *frontend.CategoryGetListCommonReq) (res *frontend.CategoryGetListCommonRes, err error) {
//	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &frontend.CategoryGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}
