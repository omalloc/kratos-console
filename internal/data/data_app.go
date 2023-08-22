package data

import (
	"context"
	"github.com/omalloc/kratos-console/internal/biz"
)

type appRepo struct {
	data *Data
}

func (r *appRepo) List(ctx context.Context, query *biz.QueryPager) ([]*biz.App, int64, error) {
	var (
		apps  []*biz.App
		count int64
	)
	err := r.data.db.WithContext(ctx).Model(&biz.App{}).
		Offset(int((query.Current - 1) * query.PageSize)).
		Limit(int(query.PageSize)).
		Count(&count).
		Find(&apps).Error
	return apps, count, err
}

func NewAppRepo(data *Data) biz.AppRepo {
	return &appRepo{
		data: data,
	}
}
