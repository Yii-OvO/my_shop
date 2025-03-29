package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"my_shop/internal/model/do"
)

type AddCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
	ParentId uint
	Content  string
}

type AddCommentOutput struct {
	Id uint
}

type DeleteCommentInput struct {
	Id uint
}

type DeleteCommentOutput struct {
	Id uint
}

// CommentListInput 获取评论列表
type CommentListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	//Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	Type uint8 // 类型
}

// CommentListOutput 查询列表结果
type CommentListOutput struct {
	List  []CommentListOutputItem
	Page  int
	Size  int
	Total int
}

type CommentListOutputItem struct {
	Id        int         `json:"id"        description:""`
	UserId    int         `json:"user_id"    description:"用户id"`
	ObjectId  int         `json:"object_id"  description:"对象id"`
	Type      int         `json:"type"      description:"评论类型：1商品 2文章"`
	ParentId  uint        `json:"parent_id" description:"父级评论id"`
	Content   string      `json:"content"   description:"评论内容"`
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CommentInfoBase struct {
	do.CommentInfo
	UserInfo UserInfoBase `json:"user" orm:"with:id=user_id"`
}
