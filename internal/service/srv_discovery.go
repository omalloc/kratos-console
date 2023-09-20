package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-agent/api/agent"
	pb "github.com/omalloc/kratos-console/api/console/discovery"
	"github.com/samber/lo"
	"strings"
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

	// 根据命名空间分组
	namespaceMap := lo.GroupBy(reply.Data, func(item *agent.Microservice) string {
		return item.Namespace
	})

	result := lo.MapToSlice(namespaceMap, func(namespace string, value []*agent.Microservice) []*pb.ServiceGroup {
		// 服务名分组
		svcMap := lo.GroupBy(value, func(item *agent.Microservice) string {
			return item.Name
		})

		return lo.MapToSlice(svcMap, func(name string, items []*agent.Microservice) *pb.ServiceGroup {
			var (
				endpoints   = make([]string, 0, len(items)*2)
				clusters    = make([]string, 0, len(items))
				seen        = make(map[string]struct{}, len(items))
				lastKey     = ""
				lastVersion = ""
			)

			for _, svc := range items {
				if lastKey == "" {
					sub := strings.LastIndex(svc.Key, "/")
					lastKey = svc.Key[0:sub]
				}
				if svc.Version != "" {
					lastVersion = svc.Version
				}
				endpoints = append(endpoints, svc.Endpoints...)
				if _, ok := seen[svc.Cluster]; ok {
					continue
				}
				seen[svc.Cluster] = struct{}{}
				clusters = append(clusters, svc.Cluster)
			}

			return &pb.ServiceGroup{
				Key:       lastKey,
				Name:      name,
				Hostname:  "",
				Namespace: namespace,
				Version:   lastVersion,
				Endpoints: endpoints,
				Clusters:  clusters,
				Children: lo.Map(items, func(item *agent.Microservice, _ int) *pb.Service {
					return &pb.Service{
						Key:       item.Key,
						Name:      item.Name,
						Hostname:  item.Id,
						Version:   item.Version,
						Endpoints: item.Endpoints,
						Cluster:   item.Cluster,
						Hang:      item.HasHang(),
						Metadata:  item.Metadata,
						Namespace: item.Namespace,
					}
				}),
			}
		})
	})

	return &pb.OnlineServiceReply{
		Pagination: nil,
		Data:       lo.Flatten(result),
	}, nil
}
