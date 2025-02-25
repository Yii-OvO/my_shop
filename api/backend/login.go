package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" tags:"登录" summary:"执行登录请求"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}
type LoginDoRes struct {
	//Info interface{} `json:"info"`
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"post" tags:"登录" summary:"刷新token"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post" tags:"登录" summary:"登出"`
}

type LogoutRes struct {
}
