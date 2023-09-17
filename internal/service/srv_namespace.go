package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/api/types"
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
	wrap := types.Wrap(req.Pagination)
	ret, err := s.usecase.List(ctx, wrap)
	if err != nil {
		return nil, err
	}

	return &pb.NamespaceListReply{
		Pagination: wrap.Resp(),
		Data: lo.Map(ret, func(item *biz.Namespace, _ int) *pb.NamespaceInfo {
			return &pb.NamespaceInfo{
				Id:          item.ID,
				Name:        item.Name,
				Alias:       item.Alias,
				Description: item.Description,
				CreatedAt:   item.CreatedAt.Format(time.DateTime),
				UpdatedAt:   item.UpdatedAt.Format(time.DateTime),
			}
		}),
	}, nil
}
func (s *NamespaceService) Get(ctx context.Context, req *pb.NamespaceGetRequest) (*pb.NamespaceGetReply, error) {
	return &pb.NamespaceGetReply{}, nil
}
func (s *NamespaceService) Create(ctx context.Context, req *pb.NamespaceCreateRequest) (*pb.NamespaceCreateReply, error) {
	return &pb.NamespaceCreateReply{}, nil
}
func (s *NamespaceService) Update(ctx context.Context, req *pb.NamespaceUpdateRequest) (*pb.NamespaceUpdateReply, error) {
	return &pb.NamespaceUpdateReply{}, nil
}
func (s *NamespaceService) Delete(ctx context.Context, req *pb.NamespaceDeleteRequest) (*pb.NamespaceDeleteReply, error) {
	return &pb.NamespaceDeleteReply{}, nil
}
