package service

import (
	"context"
	"my_shop/internal/model"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		//自定义业务参数
		Realm:           "myshop",                                           // 用户的领域名称，必传
		Key:             []byte("secret key"),                               // jwt第三部分签名所需的密钥，类型laravel中的JWT_KEY
		Timeout:         time.Minute * 60,                                   // token过期时间
		MaxRefresh:      time.Minute * 60,                                   // token过期后，可凭借旧token获取新token的刷新时间
		IdentityKey:     "id",                                               // 身份验证的key值
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token->Authorization
		TokenHeadName:   "Bearer",                                           // token在请求头时的名称
		TimeFunc:        time.Now,                                           // 设置JWT的时区
		Authenticator:   Authenticator,                                      // 根据登录信息对用户进行身份验证的回调函数，回调必须携带身份验证的键值对，如：IdentityKey:"id"，则身份认证回调函数应当包含{“id”:1}
		Unauthorized:    Unauthorized,                                       // 处理不进行授权的逻辑
		PayloadFunc:     PayloadFunc,                                        // 登录期间的设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
		IdentityHandler: IdentityHandler,                                    // 解析并设置用户身份信息，并设置身份信息至每次请求中
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}

	if user := Admin().GetAdminByNamePassword(ctx, in); user != nil {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
