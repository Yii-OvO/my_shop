package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/frontend"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

var Comment = cComment{}

type cComment struct{}

func (c *cComment) Add(ctx context.Context, req *frontend.AddCommentReq) (res *frontend.AddCommentRes, err error) {
	data := model.AddCommentInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCommentRes{Id: out.Id}, nil
}

func (a *cComment) Delete(ctx context.Context, req *frontend.DeleteCommentReq) (res *frontend.DeleteCommentRes, err error) {
	data := model.DeleteCommentInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	comment, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCommentRes{Id: comment.Id}, nil
}

func (c *cComment) List(ctx context.Context, req *frontend.CommentListReq) (res *frontend.CommentListRes, err error) {
	getListRes, err := service.Comment().GetList(ctx, model.CommentListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	return &frontend.CommentListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
