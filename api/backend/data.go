package backend

import "github.com/gogf/gf/v2/frame/g"

type DataHeadReq struct {
	g.Meta `path:"/data/head" method:"get" tags:"数据大屏" summary:"头部卡片" desc:"数据大屏的头部信息"`
}

type DataHeadRes struct {
	TodayOrderCount int `json:"today_order_count" desc:"今日订单总是"`
	DAU             int `json:"dau" desc:"今日日活"`
	ConversionRate  int `json:"conversion_rate" desc:"转化率"`
}
