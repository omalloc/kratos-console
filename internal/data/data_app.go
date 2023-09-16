package data

import (
	"context"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
)

type appRepo struct {
	data *Data
}

func (r *appRepo) List(ctx context.Context, pagination *types.Pagination) ([]*biz.App, error) {
	var (
		apps []*biz.App
	)
	err := r.data.db.WithContext(ctx).Model(&biz.App{}).
		Scopes(pagination.Paginate()).
		Count(pagination.Count()).
		Find(&apps).Error
	return apps, err
}

func NewAppRepo(data *Data) biz.AppRepo {
	return &appRepo{
		data: data,
	}
}
