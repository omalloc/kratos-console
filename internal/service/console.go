package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/omalloc/kratos-console/api/console"
)

type ConsoleService struct {
	pb.UnimplementedConsoleServer

	log *log.Helper
}

func NewConsoleService(logger log.Logger) *ConsoleService {
	return &ConsoleService{
		log: log.NewHelper(logger),
	}
}
