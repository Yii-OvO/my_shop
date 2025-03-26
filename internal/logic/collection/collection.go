package collection

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/internal/consts"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (out *model.AddCollectionOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddCollectionOutput{}, err
	}
	return &model.AddCollectionOutput{Id: uint(id)}, nil
}

func (s *sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out *model.DeleteCollectionOutput, err error) {
	// 兼容处理，优先根据收藏id删除；
	if in.Id != 0 {
		_, err = dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeleteCollectionOutput{Id: in.Id}, nil
	} else { // 收藏id为0，再根据对象id和type删除
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty(). // 注意：需要过滤空值
									Where(in).Delete()
		if err != nil {
			return &model.DeleteCollectionOutput{}, err
		}
		return &model.DeleteCollectionOutput{Id: gconv.Uint(id)}, nil
	}
}

// GetList 查询收藏列表
func (s *sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {

	var (
		m = dao.CollectionInfo.Ctx(ctx).Where(dao.CollectionInfo.Columns().UserId, ctx.Value(consts.CtxUserId)) //限制查询当前登录用户的收藏
	)
	out = &model.CollectionListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CollectionListOutputItem{}, //数据为空时返回空数组，而不是null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)

	// 条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}

	// 排序方式
	//listModel = listModel.OrderDesc(dao.CollectionInfo.Columns().Sort)
	//// 执行查询  --------------旧方法
	//var list []*entity.CollectionInfo
	//if err := listModel.WithAll().Scan(&list); err != nil {
	//	return out, err
	//}
	//// 没有数据
	//if len(list) == 0 {
	//	return out, nil
	//}
	//out.Total, err = m.Count()
	//if err != nil {
	//	return out, err
	//}

	//优化：优先查询count，而不是像之前一样查sql结果赋值到结构体中-------------------新方法
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, nil
	}

	////Collection -------------旧
	//if err := listModel.WithAll().Scan(&out.List); err != nil {
	//
	//	return out, err
	//}

	// ------------进一步优化：只查询相关的模型关联
	if in.Type == consts.CollectionTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CollectionTypeArticle {
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
