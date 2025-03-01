package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/service"
	"my_shop/utility"
)

type sData struct{}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000), // todo
		ConversionRate:  utility.RandInt(50),   // todo
	}, nil

}

// 查询今天的订单总数
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
