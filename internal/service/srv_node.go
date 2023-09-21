package service

import (
	"context"
	"github.com/omalloc/contrib/protobuf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/biz"
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
	pagination := protobuf.PageWrap(req.Pagination)
	nodes, err := s.node.GetNodeList(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return &pb.NodeListReply{
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
				ZoneId:     int32(node.ZoneID),
				ZoneName:   node.ZoneName,
				ZoneCode:   node.ZoneCode,
				RegionName: node.RegionName,
				RegionCode: node.RegionCode,
				Env:        node.Env,
				AutoDetect: node.AutoDetect,
				CreatedAt:  node.CreatedAt.Unix(),
				UpdatedAt:  node.UpdatedAt.Unix(),
			}
		}),
		Pagination: pagination.Resp(),
	}, nil
}

func (s *NodeService) Create(ctx context.Context, req *pb.CreateNodeRequest) (*pb.CreateNodeReply, error) {
	node, err := s.node.CreateNode(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNodeReply{
		Data: s.toNodeInfo(node),
	}, nil
}

func (s *NodeService) Update(ctx context.Context, req *pb.UpdateNodeRequest) (*pb.UpdateNodeReply, error) {
	node, err := s.node.UpdateNode(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateNodeReply{
		Data: s.toNodeInfo(node),
	}, nil
}

func (s *NodeService) toNodeInfo(node *biz.Node) *pb.NodeInfo {
	return &pb.NodeInfo{
		Id:         int32(node.ID),
		Name:       node.Name,
		Ip:         node.IP,
		AutoDetect: node.AutoDetect,
		Env:        node.Env,
	}
}
