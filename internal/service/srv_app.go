package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
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
				Type:      pb.AppInfo_AppType(app.Type),
				Users:     app.Users,
				Repos:     []string{},
				CreatedAt: app.CreatedAt.Unix(),
				UpdatedAt: app.UpdatedAt.Unix(),
			}
		}),
		Pagination: pagination.Resp(),
	}, nil
}
