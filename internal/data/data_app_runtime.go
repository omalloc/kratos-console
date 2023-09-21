package data

import (
	"context"
	"github.com/omalloc/contrib/kratos/orm/crud"
	"github.com/omalloc/kratos-console/internal/biz"
)

type appRuntimeRepo struct {
	crud.CRUD[biz.AppRuntime]

	data *Data
}

func (r *appRuntimeRepo) SelectListByNamespace(ctx context.Context, s string) ([]*biz.AppRuntime, error) {
	var (
		list []*biz.AppRuntime
		err  error
	)
	err = r.data.db.WithContext(ctx).
		Model(&biz.AppRuntime{}).
		Where("namespace = ?", s).
		Find(&list).Error

	return list, err
}

func NewAppRuntimeRepo(data *Data) biz.AppRuntimeRepo {
	return &appRuntimeRepo{
		CRUD: crud.New[biz.AppRuntime](data.db),
		data: data,
	}
}
