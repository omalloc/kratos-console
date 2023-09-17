package data

import (
	"context"

	"github.com/omalloc/kratos-console/internal/biz"
)

type zoneRepo struct {
	data *Data
}

func (r *zoneRepo) GetZoneList(ctx context.Context, query *biz.ZoneQuery) ([]*biz.Zone, error) {
	var zones []*biz.Zone

	tx := r.data.db.WithContext(ctx).
		Model(&biz.Zone{})
	if query.Env != "" {
		tx.Where("env = ?", query.Env)
	}

	err := tx.Scopes(query.Paginate()).
		Count(query.Count()).
		Find(&zones).
		Error

	return zones, err
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

func (r *zoneRepo) SelectSimpleList(ctx context.Context) ([]*biz.Zone, error) {
	var (
		zones []*biz.Zone
		err   error
	)

	err = r.data.db.WithContext(ctx).
		Model(&biz.Zone{}).
		Order("id DESC").
		Find(&zones).Error
	return zones, err
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
		Omit("id").
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
