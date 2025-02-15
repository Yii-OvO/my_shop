// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PraiseInfoDao is the data access object for the table praise_info.
type PraiseInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns PraiseInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PraiseInfoColumns defines and stores column names for the table praise_info.
type PraiseInfoColumns struct {
	Id        string // 点赞表
	UserId    string //
	Type      string // 点赞类型 1商品 2文章
	ObjectId  string // 点赞对象id 方便后期扩展
	CreatedAt string //
	UpdatedAt string //
}

// praiseInfoColumns holds the columns for the table praise_info.
var praiseInfoColumns = PraiseInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	Type:      "type",
	ObjectId:  "object_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPraiseInfoDao creates and returns a new DAO object for table data access.
func NewPraiseInfoDao() *PraiseInfoDao {
	return &PraiseInfoDao{
		group:   "default",
		table:   "praise_info",
		columns: praiseInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PraiseInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PraiseInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PraiseInfoDao) Columns() PraiseInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PraiseInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PraiseInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PraiseInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
