package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/model/entity"
	"my_shop/internal/service"
	"my_shop/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	// 调试打印
	gutil.Dump("加密后密码：", utility.EncryptPassword(in.Password, adminInfo.UserSalt))
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}
