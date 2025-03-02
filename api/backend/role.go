package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" tags:"角色" sm:"添加角色" dc:"添加角色"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" tags:"角色" sm:"修改角色"`
	Id     uint   `json:"id"   v:"required#角色Id不能为空" dc:"角色Id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"描述"`
}

type RoleUpdateRes struct {
	RoleId uint `json:"id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"角色" summary:"删除角色"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"角色" summary:"角色列表"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
