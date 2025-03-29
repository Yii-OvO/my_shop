package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"my_shop/internal/model/entity"
)

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"前台商品" summary:"商品列表"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"post" tags:"前台商品" summary:"商品详情"`
	Id     uint `json:"id"`
}

type GoodsDetailRes struct {
	entity.GoodsInfo
	Options   interface{} `json:"options"` // 规格sku
	Comments  interface{} `json:"comments"`
	IsCollect bool        `json:"is_collect"` // todo 需优化为单独一个接口，原因是判断是否已收藏需要获取当前登录用户id
}
