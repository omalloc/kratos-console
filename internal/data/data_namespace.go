package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/protobuf"
	"github.com/omalloc/kratos-console/internal/biz"
	"gorm.io/gorm"
)

type namespaceRepo struct {
	log  *log.Helper
	data *Data
}

func (r *namespaceRepo) SelectList(ctx context.Context, pagination *protobuf.Pagination, name string) ([]*biz.NamespaceJoinAppRuntime, error) {
	var (
		ret []*biz.NamespaceJoinAppRuntime
		err error
	)

	t2 := r.data.db.Model(&biz.AppRuntime{}).
		Select("namespace, count(id) AS running").
		Group("namespace")

	err = r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Select("namespaces.*", "IFNULL(t2.running, 0) AS running").
		Joins("LEFT JOIN (?) AS t2 ON t2.namespace = namespaces.name", t2).
		Scopes(ApplyNamespaceQuery(name)).
		Count(pagination.Count()).
		Scopes(pagination.Paginate()).
		Find(&ret).Error
	return ret, err
}

func (r *namespaceRepo) SelectSimpleAll(ctx context.Context) ([]*biz.Namespace, error) {
	var (
		ret []*biz.Namespace
		err error
	)

	err = r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Select("id, `name`, alias").
		Find(&ret).Error
	return ret, err
}

func (r *namespaceRepo) SelectOne(ctx context.Context, i int) (*biz.Namespace, error) {
	var (
		ret *biz.Namespace
		err error
	)

	err = r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Where("id = ?", i).
		Find(ret).Error
	return ret, err
}

func (r *namespaceRepo) SelectByName(ctx context.Context, s string) (*biz.Namespace, error) {
	var (
		ret *biz.Namespace
		err error
	)

	err = r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", s)).
		Find(ret).Error
	return ret, err
}

func (r *namespaceRepo) Create(ctx context.Context, namespace *biz.Namespace) error {
	return r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Create(namespace).Error
}

func (r *namespaceRepo) Update(ctx context.Context, namespace *biz.Namespace) error {
	return r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Omit("id").
		Where("id = ?", namespace.ID).
		Updates(namespace).Error
}

func (r *namespaceRepo) Delete(ctx context.Context, i int) error {
	return r.data.db.WithContext(ctx).
		Delete(&biz.Namespace{}, i).Error
}

func (r *namespaceRepo) DeleteByName(ctx context.Context, s string) error {
	return r.data.db.WithContext(ctx).
		Delete(&biz.Namespace{}, "name = ?", s).Error
}

func ApplyNamespaceQuery(name string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name != "" {
			tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
		}
		return tx
	}
}

func NewNamespaceRepo(logger log.Logger, data *Data) biz.NamespaceRepo {
	return &namespaceRepo{
		log:  log.NewHelper(logger),
		data: data,
	}
}
