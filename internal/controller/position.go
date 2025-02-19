package controller

import (
	"context"
	"my_shop/api/backend"

	"my_shop/internal/model"
	"my_shop/internal/service"
)

// 承上启下  mvc
// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	return
}

func (a *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	return &backend.PositionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

//func (a *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
//	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &backend.PositionGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}

// 前台的取值方法
//func (a *cPosition) ListFrontend(ctx context.Context, req *frontend.PositionGetListCommonReq) (res *frontend.PositionGetListCommonRes, err error) {
//	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
//		Page: req.Page,
//		Size: req.Size,
//		Sort: req.Sort,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &frontend.PositionGetListCommonRes{List: getListRes.List,
//		Page:  getListRes.Page,
//		Size:  getListRes.Size,
//		Total: getListRes.Total}, nil
//}
