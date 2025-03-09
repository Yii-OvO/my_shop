package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"my_shop/internal/model/entity"
	"time"
)

type LoginDoReq struct {
	//g.Meta   `path:"/login" method:"post" tags:"登录" summary:"执行登录请求"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// LoginDoRes for jwt
type LoginDoRes struct {
	//Info interface{} `json:"info"`
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// LoginRes for gtoken
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post" tags:"登录" summary:"刷新token"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"登录" summary:"登出"`
}

type LogoutRes struct {
}
