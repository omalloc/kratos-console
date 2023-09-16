package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/api/types"
)

type Namespace struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name;type:varchar(32);comment:命名空间唯一名称"`
	Alias       string `json:"alias" gorm:"column:alias;type:varchar(32);comment:命名空间别称"`
	Description string `json:"description" gorm:"column:description;type:varchar(128);comment:命名空间描述"`

	types.DBModel
}

func (Namespace) TableName() string {
	return "namespaces"
}

type NamespaceRepo interface {
	SelectList(context.Context, *types.Pagination) ([]*Namespace, error)
	SelectOne(context.Context, int) (*Namespace, error)
	SelectByName(context.Context, string) (*Namespace, error)
	Create(context.Context, *Namespace) error
	Update(context.Context, *Namespace) error
	Delete(context.Context, int) error
	DeleteByName(context.Context, string) error
}

type NamespaceUsecase struct {
	repo NamespaceRepo
	log  *log.Helper
}

func NewNamespaceUsecase(logger log.Logger, repo NamespaceRepo) *NamespaceUsecase {
	return &NamespaceUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (r *NamespaceUsecase) List(ctx context.Context, pagination *types.Pagination) ([]*Namespace, error) {
	return r.repo.SelectList(ctx, pagination)
}
