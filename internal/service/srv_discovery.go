package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-agent/api/agent"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/discovery"
)

type DiscoveryService struct {
	pb.UnimplementedDiscoveryServer

	log    *log.Helper
	client agent.AgentClient
}

func NewDiscoveryService(logger log.Logger, client agent.AgentClient) *DiscoveryService {
	return &DiscoveryService{
		log:    log.NewHelper(logger),
		client: client,
	}
}

func (s *DiscoveryService) OnlineServices(ctx context.Context, req *pb.OnlineServiceRequest) (*pb.OnlineServiceReply, error) {
	// 获取全部在线服务
	reply, err := s.client.ListService(ctx, &agent.ListServiceRequest{})
	if err != nil {
		return nil, err
	}

	// 根据服务名称分组
	mapData := lo.GroupBy(reply.Data, func(item *agent.Microservice) string {
		return item.Name
	})

	result := lo.MapToSlice(mapData, func(key string, value []*agent.Microservice) *pb.ServiceGroup {
		var (
			endpoints = make([]string, 0, len(value)*2)
			clusters  = make([]string, 0, len(value))
			seen      = make(map[string]struct{}, len(value))

			lastVersion string = ""
		)

		for _, service := range value {
			if service.Version != "" {
				lastVersion = service.Version
			}
			endpoints = append(endpoints, service.Endpoints...)
			if _, ok := seen[service.Cluster]; ok {
				continue
			}
			seen[service.Cluster] = struct{}{}
			clusters = append(clusters, service.Cluster)
		}

		return &pb.ServiceGroup{
			Key:       key,
			Name:      key,
			Hostname:  "",
			Version:   lastVersion,
			Endpoints: endpoints,
			Clusters:  clusters,
			Children: lo.Map(value, func(item *agent.Microservice, _ int) *pb.Service {
				hang := false
				if v, ok := item.Metadata["hang"]; ok && v == "true" {
					hang = true
				}
				return &pb.Service{
					Key:       item.Key,
					Name:      key,
					Hostname:  item.Id,
					Version:   item.Version,
					Endpoints: item.Endpoints,
					Cluster:   item.Cluster,
					Hang:      hang,
					Metadata:  item.Metadata,
				}
			}),
		}
	})

	return &pb.OnlineServiceReply{
		Pagination: nil,
		Data:       result,
	}, nil
}
