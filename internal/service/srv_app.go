package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
)

// AppService 应用服务
// 集群内应用抽象管理分组
type AppService struct {
	pb.UnimplementedAppServer

	log *log.Helper
	app *biz.AppUsecase
}

func NewAppService(logger log.Logger, app *biz.AppUsecase) *AppService {
	return &AppService{
		log: log.NewHelper(logger),
		app: app,
	}
}

func (s *AppService) List(ctx context.Context, req *pb.AppListRequest) (*pb.AppListReply, error) {
	pagination := types.Wrap(req.Pagination)
	apps, err := s.app.GetAppList(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return &pb.AppListReply{
		Data: lo.Map(apps, func(app *biz.App, _ int) *pb.AppInfo {
			return &pb.AppInfo{
				Id:          app.ID,
				Name:        app.Name,
				Alias:       app.Alias,
				Description: app.Description,
				Icon:        app.Icon,
				Ports: lo.Map(app.Ports, func(item biz.Port, _ int) *pb.AppInfo_AppPort {
					return &pb.AppInfo_AppPort{
						Port:     uint32(item.Port),
						Protocol: pb.AppInfo_AppPort_Protocol(pb.AppInfo_AppPort_Protocol_value[item.Protocol]),
						Remark:   item.Remark,
					}
				}),
				Type:        pb.AppInfo_AppType(app.Type),
				Users:       app.Users,
				Repos:       []string{},
				NamespaceId: app.NamespaceID,
				CreatedAt:   app.CreatedAt.Unix(),
				UpdatedAt:   app.UpdatedAt.Unix(),
			}
		}),
		Pagination: pagination.Resp(),
	}, nil
}

func (s *AppService) NamespaceAppList(ctx context.Context, req *pb.NamespaceAppListRequest) (*pb.NamespaceAppListReply, error) {
	ret, err := s.app.GetSimpleList(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.NamespaceAppListReply{
		Data: lo.Map(ret, func(item *biz.NamespaceApp, _ int) *pb.NamespaceApp {
			return &pb.NamespaceApp{
				Id:             item.ID,
				Name:           item.Name,
				Alias:          item.Alias,
				Type:           pb.AppInfo_AppType(item.Type),
				NamespaceId:    item.NamespaceID,
				NamespaceName:  item.NamespaceName,
				NamespaceAlias: item.NamespaceAlias,
			}
		}),
	}, nil
}

func (s *AppService) Create(ctx context.Context, req *pb.AppCreateRequest) (*pb.AppCreateReply, error) {
	data := s.toData(req.Data)
	if err := s.app.CreateApp(ctx, data); err != nil {
		return nil, err
	}

	req.Data.Id = data.ID
	return &pb.AppCreateReply{
		Data: req.Data,
	}, nil
}

func (s *AppService) toData(req *pb.AppInfo) *biz.App {
	return &biz.App{
		Name:        req.Name,
		Alias:       req.Alias,
		Description: req.Description,
		Type:        int(req.Type.Number()),
		NamespaceID: req.NamespaceId,
		Icon:        req.Icon,
		Users:       req.Users,
		Repos:       req.Repos,
		Ports: lo.Map(req.Ports, func(item *pb.AppInfo_AppPort, _ int) biz.Port {
			return biz.Port{
				Port:     uint(item.Port),
				Protocol: item.Protocol.String(),
				Remark:   item.Remark,
			}
		}),
	}
}
