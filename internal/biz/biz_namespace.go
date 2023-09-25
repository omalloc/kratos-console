package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/kratos/orm"
	"github.com/omalloc/contrib/protobuf"
)

var (
	ErrInvalidNamespaceID = errors.New(400001, "InvalidNamespaceID", "无效的命名空间")
)

type Namespace struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name;type:varchar(32);comment:命名空间唯一名称"`
	Alias       string `json:"alias" gorm:"column:alias;type:varchar(32);comment:命名空间别称"`
	Description string `json:"description" gorm:"column:description;type:varchar(128);comment:命名空间描述"`

	orm.DBModel
}

func (Namespace) TableName() string {
	return "namespaces"
}

type NamespaceJoinAppRuntime struct {
	Namespace
	Running int64 `json:"running" gorm:"column:running;type:int;comment:运行中的应用数量"`
}

type NamespaceRepo interface {
	SelectList(context.Context, *protobuf.Pagination, string) ([]*NamespaceJoinAppRuntime, error)
	SelectSimpleAll(context.Context) ([]*Namespace, error)
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

func (r *NamespaceUsecase) List(ctx context.Context, pagination *protobuf.Pagination, name string) ([]*NamespaceJoinAppRuntime, error) {
	return r.repo.SelectList(ctx, pagination, name)
}

func (r *NamespaceUsecase) SimpleList(ctx context.Context) ([]*Namespace, error) {
	return r.repo.SelectSimpleAll(ctx)
}

func (r *NamespaceUsecase) Create(ctx context.Context, namespace *Namespace) error {
	return r.repo.Create(ctx, namespace)
}

func (r *NamespaceUsecase) Update(ctx context.Context, namespace *Namespace) error {
	if namespace.ID <= 0 {
		return ErrInvalidNamespaceID
	}
	return r.repo.Update(ctx, namespace)
}

func (r *NamespaceUsecase) Delete(ctx context.Context, id int64) error {
	return r.repo.Delete(ctx, int(id))
}
