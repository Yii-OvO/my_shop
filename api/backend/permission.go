package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionReq struct {
	g.Meta `path:"/permission/add" method:"post" tags:"权限" sm:"添加权限" dc:"添加权限"`
	PermissionCreatUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update" method:"post" tags:"权限" sm:"修改权限"`
	Id     uint `json:"id"   v:"required#权限Id不能为空" dc:"权限Id"`
	PermissionCreatUpdateBase
}

type PermissionCreatUpdateBase struct {
	Name string `json:"name" v:"required#权限名称不能为空" dc:"权限名称"`
	Path string `json:"path" v:"required#权限路径不能为空" dc:"权限路径"`
}

type PermissionUpdateRes struct {
	PermissionId uint `json:"id"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete" method:"delete" tags:"权限" summary:"删除权限"`
	Id     uint `v:"min:1#请选择需要删除的权限" dc:"权限id"`
}
type PermissionDeleteRes struct{}

type PermissionGetListCommonReq struct {
	g.Meta `path:"/permission/list" method:"get" tags:"权限" summary:"权限列表"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
