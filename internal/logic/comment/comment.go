package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"my_shop/internal/consts"
	"my_shop/internal/dao"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) AddComment(ctx context.Context, in model.AddCommentInput) (out *model.AddCommentOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddCommentOutput{}, err
	}
	return &model.AddCommentOutput{Id: uint(id)}, nil
}

func (s *sComment) DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out *model.DeleteCommentOutput, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	}
	_, err = dao.CommentInfo.Ctx(ctx).Where(condition).Delete()
	if err != nil {
		return nil, err
	}
	return &model.DeleteCommentOutput{Id: in.Id}, nil
}

// GetList 查询评论列表
func (s *sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {

	var (
		m = dao.CommentInfo.Ctx(ctx).Where(dao.CommentInfo.Columns().UserId, ctx.Value(consts.CtxUserId)) //限制查询当前登录用户的评论
	)
	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CommentListOutputItem{}, //数据为空时返回空数组，而不是null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)

	// 条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
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
	if in.Type == consts.CommentTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
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

// CommentCount 抽取获得评论数量的方法 for 商品详情&文章详情
func CommentCount(ctx context.Context, objectId uint, commentType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     commentType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CheckComment 抽取判断当前用户是否评论方法 for 商品详情&文章详情
func CheckComment(ctx context.Context, in model.CheckCollectionInput) (bool, error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count >= 0 {
		return true, nil
	} else {
		return false, nil
	}
}
