package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"my_shop/internal/consts"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/model/do"
	"my_shop/internal/service"
	"my_shop/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {

	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回Id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// UpdatePassword 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	// 验证密保问题
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		g.Dump("userInfo.SecretAnswer:", userInfo.SecretAnswer)
		g.Dump("in.SecretAnswer:", in.SecretAnswer)
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}
	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, err
}
