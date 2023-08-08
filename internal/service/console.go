package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
	"time"

	pb "github.com/omalloc/kratos-console/api/console"
	resource "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/biz"
)

type ConsoleService struct {
	pb.UnimplementedConsoleServer

	log  *log.Helper
	zone *biz.ZoneUsecase
}

func NewConsoleService(zone *biz.ZoneUsecase, logger log.Logger) *ConsoleService {
	return &ConsoleService{
		log:  log.NewHelper(logger),
		zone: zone,
	}
}

func (s *ConsoleService) GetZoneList(ctx context.Context, req *resource.GetZoneListRequest) (*resource.GetZoneListReply, error) {
	data, total, err := s.zone.GetZoneList(ctx, &biz.QueryPager{
		PageSize: req.PageSize,
		Current:  req.Current,
	})

	if err != nil {
		return nil, err
	}

	return &resource.GetZoneListReply{
		Data: lo.Map(data, func(item *biz.Zone, index int) *resource.Zone {
			return &resource.Zone{
				Id:         int32(item.ID),
				Name:       item.Name,
				Code:       item.Code,
				RegionName: item.RegionName,
				RegionCode: item.RegionCode,
				Env:        item.Env,
				CreatedAt:  item.CreatedAt.Format(time.RFC3339),
				UpdatedAt:  item.UpdatedAt.Format(time.RFC3339),
			}
		}),
		Total: int32(total),
	}, nil
}

func (s *ConsoleService) GetZone(ctx context.Context, req *resource.GetZoneRequest) (*resource.GetZoneReply, error) {
	return &resource.GetZoneReply{}, nil
}
func (s *ConsoleService) CreateZone(ctx context.Context, req *resource.CreateZoneRequest) (*resource.CreateZoneReply, error) {
	return &resource.CreateZoneReply{}, nil
}
func (s *ConsoleService) UpdateZone(ctx context.Context, req *resource.UpdateZoneRequest) (*resource.UpdateZoneReply, error) {
	return &resource.UpdateZoneReply{}, nil
}
func (s *ConsoleService) DisableZone(ctx context.Context, req *resource.DisableZoneRequest) (*resource.DisableZoneReply, error) {
	return &resource.DisableZoneReply{}, nil
}
func (s *ConsoleService) DeleteZone(ctx context.Context, req *resource.DeleteZoneRequest) (*resource.DeleteZoneReply, error) {
	return &resource.DeleteZoneReply{}, nil
}
