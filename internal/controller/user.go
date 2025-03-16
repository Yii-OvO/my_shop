package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/api/frontend"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{Id: out.Id}, nil
}
