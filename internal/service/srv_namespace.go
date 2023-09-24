package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/protobuf"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/biz"
)

type NamespaceService struct {
	pb.UnimplementedNamespaceServer

	log     *log.Helper
	usecase *biz.NamespaceUsecase
}

func NewNamespaceService(logger log.Logger, usecase *biz.NamespaceUsecase) *NamespaceService {
	return &NamespaceService{
		log:     log.NewHelper(logger),
		usecase: usecase,
	}
}

func (s *NamespaceService) List(ctx context.Context, req *pb.NamespaceListRequest) (*pb.NamespaceListReply, error) {
	wrap := protobuf.PageWrap(req.Pagination)
	ret, err := s.usecase.List(ctx, wrap, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.NamespaceListReply{
		Pagination: wrap.Resp(),
		Data: lo.Map(ret, func(item *biz.NamespaceJoinAppRuntime, _ int) *pb.NamespaceInfo {
			return &pb.NamespaceInfo{
				Id:          item.ID,
				Name:        item.Name,
				Alias:       item.Alias,
				Description: item.Description,
				Running:     int32(item.Running),
				CreatedAt:   item.CreatedAt.Format(time.DateTime),
				UpdatedAt:   item.UpdatedAt.Format(time.DateTime),
			}
		}),
	}, nil
}

func (s *NamespaceService) SimpleList(ctx context.Context, _ *pb.NamespaceSimpleListRequest) (*pb.NamespaceSimpleListReply, error) {
	data, err := s.usecase.SimpleList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.NamespaceSimpleListReply{
		Data: lo.Map(data, func(item *biz.Namespace, _ int) *pb.NamespaceSimple {
			return &pb.NamespaceSimple{
				Id:    item.ID,
				Name:  item.Name,
				Alias: item.Alias,
			}
		}),
	}, nil
}

func (s *NamespaceService) Get(ctx context.Context, req *pb.NamespaceGetRequest) (*pb.NamespaceGetReply, error) {
	return &pb.NamespaceGetReply{}, nil
}
func (s *NamespaceService) Create(ctx context.Context, req *pb.NamespaceCreateRequest) (*pb.NamespaceCreateReply, error) {
	m := &biz.Namespace{
		Name:        req.Name,
		Alias:       req.Alias,
		Description: req.Description,
	}
	err := s.usecase.Create(ctx, m)
	return &pb.NamespaceCreateReply{
		Data: s.toData(m),
	}, err
}
func (s *NamespaceService) Update(ctx context.Context, req *pb.NamespaceUpdateRequest) (*pb.NamespaceUpdateReply, error) {
	m := &biz.Namespace{
		ID:          req.Id,
		Name:        req.Name,
		Alias:       req.Alias,
		Description: req.Description,
	}
	err := s.usecase.Update(ctx, m)

	return &pb.NamespaceUpdateReply{
		Data: s.toData(m),
	}, err
}
func (s *NamespaceService) Delete(ctx context.Context, req *pb.NamespaceDeleteRequest) (*pb.NamespaceDeleteReply, error) {
	err := s.usecase.Delete(ctx, req.Id)
	return &pb.NamespaceDeleteReply{}, err
}

func (s *NamespaceService) toData(m *biz.Namespace) *pb.NamespaceInfo {
	return &pb.NamespaceInfo{
		Id:          m.ID,
		Name:        m.Name,
		Alias:       m.Alias,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.Format(time.DateTime),
		UpdatedAt:   m.UpdatedAt.Format(time.DateTime),
	}
}
