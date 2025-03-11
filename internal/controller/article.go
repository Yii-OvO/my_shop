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
// Article 文章

var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleReq) (res *backend.ArticleRes, err error) {
	data := model.ArticleCreateInput{}
	err = gconv.Scan(req, &data) //自动识别类型转换
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId)) //获取当前登录的管理员id
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleRes{Id: out.Id}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId)) //获取当前登录的管理员id
	err = service.Article().Update(ctx, data)
	return &backend.ArticleUpdateRes{Id: req.Id}, nil
}

func (a *cArticle) List(ctx context.Context, req *backend.ArticleGetListCommonReq) (res *backend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
