package controller

import (
	"context"
	"my_shop/api/backend"
	"my_shop/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

//func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	err = service.Login().Login(ctx, model.UserLoginInput{
//		Name:     req.Name,
//		Password: req.Password,
//	})
//	if err != nil {
//		return
//	}
//	// 识别并跳转到登录前页面
//	//res.Info = service.Session().GetUser(ctx)
//	return
//}

//for jwt
//func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
//	return
//}

// RefreshToken for jwt, 在gtoken中不需要客户端主动刷新token，而是在服务端自动刷新
func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
