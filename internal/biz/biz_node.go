package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Node struct {
	ID     int64  `json:"id" gorm:"column:id;primaryKey"`
	Name   string `json:"name" gorm:"column:name;type:varchar(64);"`
	IP     string `json:"ip" gorm:"column:ip;type:varchar(64);"`
	ZoneID int64  `json:"zone_id" gorm:"column:zone_id;type:bigint(20);"`
	Env    string `json:"env" gorm:"column:env;type:varchar(12);"`

	DBModel
}

type NodeInfo struct {
	Node

	ZoneName   string `json:"zone_name" gorm:"column:zone_name;"`
	ZoneCode   string `json:"zone_code" gorm:"column:zone_code;"`
	RegionName string `json:"region_name" gorm:"column:region_name;"`
	RegionCode string `json:"region_code" gorm:"column:region_code;"`
}

type NodeRepo interface {
	List(ctx context.Context, query *QueryPager) ([]*NodeInfo, int64, error)
}

type NodeUsecase struct {
	log  *log.Helper
	repo NodeRepo
}

func NewNodeUsecase(logger log.Logger, repo NodeRepo) *NodeUsecase {
	return &NodeUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (uc *NodeUsecase) GetNodeList(ctx context.Context, query *QueryPager) ([]*NodeInfo, int64, error) {
	return uc.repo.List(ctx, query)
}
