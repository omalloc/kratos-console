package data

import (
	"context"
	"github.com/omalloc/contrib/protobuf"
	"github.com/omalloc/kratos-console/internal/biz"
)

type nodeRepo struct {
	data *Data
}

func (r *nodeRepo) Create(ctx context.Context, node *biz.Node) error {
	return r.data.db.WithContext(ctx).
		Model(&biz.Node{}).
		Create(node).
		Error
}

func (r *nodeRepo) Update(ctx context.Context, node *biz.Node) error {
	if node.ID <= 0 {
		return biz.ErrPrimaryIDNotExist
	}
	return r.data.db.WithContext(ctx).
		Model(&biz.Node{}).
		Omit("id", "created_at").
		Where("id = ?", node.ID).
		Save(node).
		Error
}

func (r *nodeRepo) List(ctx context.Context, pagination *protobuf.Pagination) ([]*biz.NodeInfo, error) {
	var nodes []*biz.NodeInfo
	err := r.data.db.WithContext(ctx).
		Model(&biz.Node{}).
		Select("nodes.*",
			"zones.name AS zone_name",
			"zones.code AS zone_code",
			"zones.region_name AS region_name",
			"zones.region_code AS region_code").
		Joins("LEFT JOIN zones ON nodes.zone_id = zones.id").
		Scopes(pagination.Paginate()).
		Count(pagination.Count()).
		Find(&nodes).
		Error
	return nodes, err
}

func NewNodeRepo(data *Data) biz.NodeRepo {
	return &nodeRepo{
		data: data,
	}
}
