package data

import (
	"context"
	"github.com/omalloc/kratos-console/internal/biz"
)

type zoneRepo struct {
	data *Data
}

func (r *zoneRepo) GetZoneList(ctx context.Context) ([]*biz.Zone, int64, error) {
	var (
		count int64
		zones []*biz.Zone
	)

	err := r.data.db.Model(&biz.Zone{}).Count(&count).Find(&zones).Error
	return zones, count, err
}

func NewZoneRepo(data *Data) biz.ZoneRepo {
	return &zoneRepo{
		data: data,
	}
}
