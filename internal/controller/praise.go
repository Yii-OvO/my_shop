package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/frontend"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

var Praise = cPraise{}

type cPraise struct{}

func (c *cPraise) Add(ctx context.Context, req *frontend.AddPraiseReq) (res *frontend.AddPraiseRes, err error) {
	data := model.AddPraiseInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddPraiseRes{Id: out.Id}, nil
}

func (a *cPraise) Delete(ctx context.Context, req *frontend.DeletePraiseReq) (res *frontend.DeletePraiseRes, err error) {
	data := model.DeletePraiseInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	praise, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeletePraiseRes{Id: praise.Id}, nil
}

func (c *cPraise) List(ctx context.Context, req *frontend.PraiseListReq) (res *frontend.PraiseListRes, err error) {
	getListRes, err := service.Praise().GetList(ctx, model.PraiseListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	return &frontend.PraiseListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
