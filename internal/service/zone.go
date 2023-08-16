package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/samber/lo"
	"time"
)

type ZoneService struct {
	pb.UnimplementedZoneServer

	log  *log.Helper
	zone *biz.ZoneUsecase
}

func NewZoneService(logger log.Logger, zone *biz.ZoneUsecase) *ZoneService {

	return &ZoneService{
		log:  log.NewHelper(logger),
		zone: zone,
	}
}

func (s *ZoneService) GetZoneList(ctx context.Context, req *pb.GetZoneListRequest) (*pb.GetZoneListReply, error) {
	data, total, err := s.zone.GetZoneList(ctx, &biz.QueryPager{
		PageSize: req.PageSize,
		Current:  req.Current,
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetZoneListReply{
		Data: lo.Map(data, func(item *biz.Zone, index int) *pb.ZoneInfo {
			return &pb.ZoneInfo{
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

func (s *ZoneService) GetZone(ctx context.Context, req *pb.GetZoneRequest) (*pb.GetZoneReply, error) {
	return &pb.GetZoneReply{}, nil
}

func (s *ZoneService) CreateZone(ctx context.Context, req *pb.CreateZoneRequest) (*pb.CreateZoneReply, error) {
	zone, err := s.zone.CreateZone(ctx, req)
	return &pb.CreateZoneReply{
		Data: &pb.ZoneInfo{
			Id:         int32(zone.ID),
			Name:       zone.Name,
			Code:       zone.Code,
			RegionName: zone.RegionName,
			RegionCode: zone.RegionCode,
			Env:        zone.Env,
		},
	}, err
}

func (s *ZoneService) UpdateZone(ctx context.Context, req *pb.UpdateZoneRequest) (*pb.UpdateZoneReply, error) {
	return &pb.UpdateZoneReply{}, nil
}

func (s *ZoneService) DisableZone(ctx context.Context, req *pb.DisableZoneRequest) (*pb.DisableZoneReply, error) {
	return &pb.DisableZoneReply{}, nil
}

func (s *ZoneService) DeleteZone(ctx context.Context, req *pb.DeleteZoneRequest) (*pb.DeleteZoneReply, error) {
	return &pb.DeleteZoneReply{}, nil
}
