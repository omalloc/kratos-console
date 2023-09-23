package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-agent/api/agent"
	pb "github.com/omalloc/kratos-console/api/console/discovery"
	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/samber/lo"
	"strings"
)

type DiscoveryService struct {
	pb.UnimplementedDiscoveryServer

	log        *log.Helper
	client     agent.AgentClient
	appruntime *biz.AppRuntimeUsecase
}

func NewDiscoveryService(logger log.Logger, client agent.AgentClient, appruntime *biz.AppRuntimeUsecase) *DiscoveryService {
	return &DiscoveryService{
		log:        log.NewHelper(logger),
		client:     client,
		appruntime: appruntime,
	}
}

func (s *DiscoveryService) OnlineServices(ctx context.Context, req *pb.OnlineServiceRequest) (*pb.OnlineServiceReply, error) {
	ret, err := s.appruntime.SelectAll(ctx, req.Name, "")
	if err != nil {
		return nil, err
	}

	// 索引，快速寻找到对应的服务
	keyMap := lo.GroupBy(ret.Children, func(item *biz.AppRuntime) string {
		return strings.Replace(item.Key, fmt.Sprintf("/%s", item.Hostname), "", 1)
	})

	data := lo.Map(ret.MainData, func(item *biz.AppRuntime, index int) *pb.ServiceGroup {
		return &pb.ServiceGroup{
			Key:       item.Key,
			Name:      item.Name,
			Namespace: item.Namespace,
			Hostname:  "", // always empty
			Version:   item.Version,
			Clusters:  strings.Split(item.Cluster, ","),
			Endpoints: nil,
			IsGroup:   true,
			Children: lo.Map(keyMap[item.Key], func(child *biz.AppRuntime, _ int) *pb.Service {
				return &pb.Service{
					Key:       child.Key,
					Name:      child.Name,
					Hostname:  child.Hostname,
					Version:   child.Version,
					Endpoints: child.Endpoints,
					Cluster:   child.Cluster,
					Hang:      child.Metadata["hang"] == "true",
					Metadata:  child.Metadata,
					Namespace: child.Namespace,
				}
			}),
		}
	})

	return &pb.OnlineServiceReply{
		Pagination: nil,
		Data:       data,
	}, err

	/*

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
		}, nil*/
}
