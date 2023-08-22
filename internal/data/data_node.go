package data

import (
	"context"
	"github.com/omalloc/kratos-console/internal/biz"
)

type nodeRepo struct {
	data *Data
}

func (r *nodeRepo) List(ctx context.Context, query *biz.QueryPager) ([]*biz.NodeInfo, int64, error) {
	var (
		count int64
		nodes []*biz.NodeInfo
	)
	err := r.data.db.WithContext(ctx).
		Model(&biz.Node{}).
		Select("nodes.*, zones.name AS zone_name, zones.code AS zone_code, zones.region_name AS region_name, zones.region_code AS region_code").
		Joins("LEFT JOIN zones ON nodes.zone_id = zones.id").
		Offset(int((query.Current - 1) * query.PageSize)).
		Limit(int(query.PageSize)).
		Count(&count).Find(&nodes).
		Error
	return nodes, count, err
}

func NewNodeRepo(data *Data) biz.NodeRepo {
	return &nodeRepo{
		data: data,
	}
}
