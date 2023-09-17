package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
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

func (s *ZoneService) List(ctx context.Context, req *pb.GetZoneListRequest) (*pb.GetZoneListReply, error) {
	pagination := types.Wrap(req.Pagination)
	data, err := s.zone.GetZoneList(ctx, &biz.ZoneQuery{
		Pagination: pagination,
		Env:        req.Env,
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
		Pagination: pagination.Resp(),
	}, nil
}

func (s *ZoneService) SimpleList(ctx context.Context, req *pb.SimpleListRequest) (*pb.SimpleListReply, error) {
	ret, err := s.zone.GetSimpleList(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.SimpleListReply{
		Data: lo.Map(ret, func(item *biz.Zone, _ int) *pb.SimpleZoneInfo {
			return &pb.SimpleZoneInfo{
				Id:   int32(item.ID),
				Name: item.Name,
				Code: item.Code,
			}
		}),
	}, nil
}

func (s *ZoneService) Get(ctx context.Context, req *pb.GetZoneRequest) (*pb.GetZoneReply, error) {
	zone, err := s.zone.GetZone(ctx, int64(req.Id))
	if err != nil {
		return nil, errors.NotFound("ZONE_NOT_FOUND", "Zone Not Found")
	}
	return &pb.GetZoneReply{
		Data: &pb.ZoneInfo{
			Id:         int32(zone.ID),
			Name:       zone.Name,
			Code:       zone.Code,
			RegionName: zone.RegionName,
			RegionCode: zone.RegionCode,
			Env:        zone.Env,
			CreatedAt:  zone.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  zone.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *ZoneService) Create(ctx context.Context, req *pb.CreateZoneRequest) (*pb.CreateZoneReply, error) {
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

func (s *ZoneService) Update(ctx context.Context, req *pb.UpdateZoneRequest) (*pb.UpdateZoneReply, error) {
	if err := s.zone.UpdateZone(ctx, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ZoneService) Disable(ctx context.Context, req *pb.DisableZoneRequest) (*pb.DisableZoneReply, error) {
	return &pb.DisableZoneReply{}, nil
}

func (s *ZoneService) Delete(ctx context.Context, req *pb.DeleteZoneRequest) (*pb.DeleteZoneReply, error) {
	err := s.zone.DeleteZone(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteZoneReply{}, nil
}
