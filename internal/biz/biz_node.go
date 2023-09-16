package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/api/types"
)

var (
	ErrPrimaryIDNotExist = errors.New("primary id not exist")
)

type Node struct {
	ID         int    `json:"id" gorm:"column:id;primaryKey"`
	Name       string `json:"name" gorm:"column:name;type:varchar(64);"`
	IP         string `json:"ip" gorm:"column:ip;type:varchar(64);"`
	ZoneID     int    `json:"zone_id" gorm:"column:zone_id;type:bigint(20);"`
	Env        string `json:"env" gorm:"column:env;type:varchar(12);"`
	AutoDetect bool   `json:"auto_detect" gorm:"column:auto_detect;type:tinyint(1);default:0;"`

	types.DBModel
}

type NodeInfo struct {
	Node

	ZoneName   string `json:"zone_name" gorm:"column:zone_name;"`
	ZoneCode   string `json:"zone_code" gorm:"column:zone_code;"`
	RegionName string `json:"region_name" gorm:"column:region_name;"`
	RegionCode string `json:"region_code" gorm:"column:region_code;"`
	//Env        string `json:"env" gorm:"column:env"`
}

type NodeRepo interface {
	List(ctx context.Context, pagination *types.Pagination) ([]*NodeInfo, error)
	Create(ctx context.Context, node *Node) error
	Update(ctx context.Context, node *Node) error
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

func (uc *NodeUsecase) GetNodeList(ctx context.Context, pagination *types.Pagination) ([]*NodeInfo, error) {
	return uc.repo.List(ctx, pagination)
}

func (uc *NodeUsecase) CreateNode(ctx context.Context, data *pb.CreateNodeRequest) (*Node, error) {
	node := &Node{
		Name:       data.Name,
		IP:         data.Ip,
		Env:        data.Env,
		ZoneID:     int(data.ZoneId),
		AutoDetect: data.AutoDetect,
	}
	if err := uc.repo.Create(ctx, node); err != nil {
		return nil, err
	}
	return node, nil
}

func (uc *NodeUsecase) UpdateNode(ctx context.Context, data *pb.UpdateNodeRequest) (*Node, error) {
	node := &Node{
		ID:         int(data.Id),
		Name:       data.Name,
		IP:         data.Ip,
		Env:        data.Env,
		ZoneID:     int(data.ZoneId),
		AutoDetect: data.AutoDetect,
	}
	if err := uc.repo.Update(ctx, node); err != nil {
		return nil, err
	}
	return node, nil
}
