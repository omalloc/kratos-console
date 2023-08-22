package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	apps, total, err := s.app.GetAppList(ctx, &biz.QueryPager{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AppListReply{
		Total: total,
		Data: lo.Map(apps, func(app *biz.App, _ int) *pb.AppInfo {
			return &pb.AppInfo{
				Id:          app.ID,
				Name:        app.Name,
				Description: app.Description,
				Icon:        app.Icon,
				CreatedAt:   app.CreatedAt.Unix(),
				UpdatedAt:   app.UpdatedAt.Unix(),
			}
		}),
	}, nil
}
