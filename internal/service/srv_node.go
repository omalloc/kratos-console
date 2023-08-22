package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/samber/lo"
	"time"

	pb "github.com/omalloc/kratos-console/api/console/resource"
)

type NodeService struct {
	pb.UnimplementedNodeServer

	log  *log.Helper
	node *biz.NodeUsecase
}

func NewNodeService(logger log.Logger, node *biz.NodeUsecase) *NodeService {
	return &NodeService{
		log:  log.NewHelper(logger),
		node: node,
	}
}

func (s *NodeService) List(ctx context.Context, req *pb.NodeListRequest) (*pb.NodeListReply, error) {
	nodes, total, err := s.node.GetNodeList(ctx, &biz.QueryPager{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &pb.NodeListReply{
		Total: total,
		Data: lo.Map(nodes, func(node *biz.NodeInfo, _ int) *pb.NodeInfo {
			return &pb.NodeInfo{
				Id:   int32(node.ID),
				Name: node.Name,
				Ip:   node.IP,
				Agent: &pb.Agent{
					Type:          pb.AgentType_ABNORMAL,
					HeartbeatTime: time.Now().Unix(),
					Version:       "v1.0.2",
				},
				ZoneId:     node.ZoneID,
				ZoneName:   node.ZoneName,
				ZoneCode:   node.ZoneCode,
				RegionName: node.RegionName,
				RegionCode: node.RegionCode,
				CreatedAt:  node.CreatedAt.Unix(),
				UpdatedAt:  node.UpdatedAt.Unix(),
			}
		}),
	}, nil
}
