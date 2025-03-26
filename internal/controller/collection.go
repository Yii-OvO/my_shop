package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/frontend"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

var Collection = cCollection{}

type cCollection struct{}

func (c *cCollection) Add(ctx context.Context, req *frontend.AddCollectionReq) (res *frontend.AddCollectionRes, err error) {
	data := model.AddCollectionInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCollectionRes{Id: out.Id}, nil
}

func (a *cCollection) Delete(ctx context.Context, req *frontend.DeleteCollectionReq) (res *frontend.DeleteCollectionRes, err error) {
	data := model.DeleteCollectionInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCollectionRes{Id: collection.Id}, nil
}

func (c *cCollection) List(ctx context.Context, req *frontend.CollectionListReq) (res *frontend.CollectionListRes, err error) {
	getListRes, err := service.Collection().GetList(ctx, model.CollectionListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	return &frontend.CollectionListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
