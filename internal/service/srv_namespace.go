package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/api/types"
	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/samber/lo"
	"time"

	pb "github.com/omalloc/kratos-console/api/console/resource"
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

func (s *NamespaceService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	wrap := types.Wrap(req.Pagination)
	ret, err := s.usecase.List(ctx, wrap)
	if err != nil {
		return nil, err
	}

	return &pb.ListReply{
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
func (s *NamespaceService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	return &pb.GetReply{}, nil
}
func (s *NamespaceService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{}, nil
}
func (s *NamespaceService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	return &pb.UpdateReply{}, nil
}
func (s *NamespaceService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
