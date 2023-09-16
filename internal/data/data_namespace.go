package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
)

type namespaceRepo struct {
	log  *log.Helper
	data *Data
}

func (r *namespaceRepo) SelectList(ctx context.Context, pagination *types.Pagination) ([]*biz.Namespace, error) {
	var (
		ret []*biz.Namespace
		err error
	)

	err = r.data.db.WithContext(ctx).Model(&biz.Namespace{}).
		Scopes(pagination.Paginate()).
		Count(pagination.Count()).
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

func NewNamespaceRepo(logger log.Logger, data *Data) biz.NamespaceRepo {
	return &namespaceRepo{
		log:  log.NewHelper(logger),
		data: data,
	}
}
