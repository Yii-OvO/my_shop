package consts

// 当常量在多个地方需要使用，最好抽取出来
const (
	ProjectName              = "Go开源MyShop"
	ProjectUsage             = "MyShop"
	ProjectBrief             = "start http server"
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GTokenAdminPrefix        = "Admin:"             // gtoken登录 管理后台前缀区分
	GTokenFrontendPrefix     = "User:"              // gtoken登录 前台用户前缀区分
	// for backend admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminRoleIds = "CtxAdminRoleIds"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	// for frontend user
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserSex    = "CtxUserSex"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"
	// for 登录相关
	TokenType          = "Bearer"
	CacheModeRedis     = 2
	BackendServerName  = "myshop"
	MultiLogin         = true
	FrontendMultiLogin = true
	GTokenExpireIn     = 10 * 24 * 60 * 60
	// 统一管理错误提升
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFailMsg         = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg      = "密保问题不正确"
	// 收藏相关
	CollectionTypeGoods   = 1
	CollectionTypeArticle = 2
)
