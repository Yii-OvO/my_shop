package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddCommentReq struct {
	g.Meta   `path:"/comment/add" method:"post" tags:"前台评论" summary:"添加评论"`
	ObjectId uint   `json:"object_id"  description:"对象id"              v:"required#评论的对象id必填"`
	Type     uint8  `json:"type"      description:"评论类型：1商品 2文章"  v:"in:1,2"` //数据校验 范围约束
	ParentId uint   `json:"parent_id" description:"父级评论id"`
	Content  string `json:"content"   description:"评论内容"  v:"required#评论内容不能为空"`
}

type AddCommentRes struct {
	Id uint `json:"id"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/comment/delete" method:"post" tags:"前台评论" summary:"删除评论"`
	Id     uint `json:"id"`
}

type DeleteCommentRes struct {
	Id uint `json:"id"`
}

type CommentListReq struct {
	g.Meta `path:"/comment/list" method:"post" tags:"前台评论" summary:"评论列表"`
	Type   uint8 `json:"type"      description:"评论类型"  v:"in:0,1,2"` // 类型
	CommonPaginationReq
}

type CommentListRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListCommentItem struct {
	Id       int         `json:"id"        description:""`
	UserId   int         `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"object_id"  description:"对象id"`
	Type     int         `json:"type"      description:"评论类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
