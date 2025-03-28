package praise

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/internal/consts"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) AddPraise(ctx context.Context, in model.AddPraiseInput) (out *model.AddPraiseOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddPraiseOutput{}, err
	}
	return &model.AddPraiseOutput{Id: uint(id)}, nil
}

func (s *sPraise) DeletePraise(ctx context.Context, in model.DeletePraiseInput) (out *model.DeletePraiseOutput, err error) {
	// 兼容处理，优先根据点赞id删除；
	if in.Id != 0 {
		_, err = dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeletePraiseOutput{Id: in.Id}, nil
	} else { // 点赞id为0，再根据对象id和type删除
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty(). // 注意：需要过滤空值
								Where(in).Delete()
		if err != nil {
			return &model.DeletePraiseOutput{}, err
		}
		return &model.DeletePraiseOutput{Id: gconv.Uint(id)}, nil
	}
}

// GetList 查询点赞列表
func (s *sPraise) GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {

	var (
		m = dao.PraiseInfo.Ctx(ctx).Where(dao.PraiseInfo.Columns().UserId, ctx.Value(consts.CtxUserId)) //限制查询当前登录用户的点赞
	)
	out = &model.PraiseListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.PraiseListOutputItem{}, //数据为空时返回空数组，而不是null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)

	// 条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}

	//优化：优先查询count，而不是像之前一样查sql结果赋值到结构体中-------------------新方法
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, nil
	}

	// ------------进一步优化：只查询相关的模型关联
	if in.Type == consts.PraiseTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}

// PraiseCount 抽取获得点赞数量的方法 for 商品详情&文章详情
func PraiseCount(ctx context.Context, objectId uint, praiseType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     praiseType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CheckPraise 抽取判断当前用户是否点赞方法 for 商品详情&文章详情
func CheckPraise(ctx context.Context, in model.CheckCollectionInput) (bool, error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count >= 0 {
		return true, nil
	} else {
		return false, nil
	}
}
