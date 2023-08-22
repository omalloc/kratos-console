package data

import (
	"context"
	"github.com/omalloc/kratos-console/internal/biz"
)

type zoneRepo struct {
	data *Data
}

func (r *zoneRepo) GetZoneList(ctx context.Context, query *biz.QueryPager) ([]*biz.Zone, int64, error) {
	var (
		count int64
		zones []*biz.Zone
	)

	err := r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Count(&count).
		Offset(int((query.Current - 1) * query.PageSize)).
		Limit(int(query.PageSize)).
		Find(&zones).Error
	return zones, count, err
}

func (r *zoneRepo) GetZoneByID(ctx context.Context, ID int64) (*biz.Zone, error) {
	var (
		zone biz.Zone
		err  error
	)

	err = r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Where("id = ?", ID).
		Take(&zone).
		Error
	return &zone, err
}

func (r *zoneRepo) CreateZone(ctx context.Context, zone *biz.Zone) error {
	return r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Create(zone).
		Error
}

func (r *zoneRepo) UpdateZone(ctx context.Context, zone *biz.Zone) error {
	return r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Where("id = ?", zone.ID).
		Updates(zone).
		Error
}

func (r *zoneRepo) DisableZone(ctx context.Context, ID int64) error {
	return r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Where("id = ?", ID).
		Update("status", 2).
		Error
}

func (r *zoneRepo) DeleteZone(ctx context.Context, ID int64) error {
	return r.data.db.WithContext(ctx).
		Delete(&biz.Zone{}, ID).
		Error
}

func NewZoneRepo(data *Data) biz.ZoneRepo {
	return &zoneRepo{
		data: data,
	}
}
