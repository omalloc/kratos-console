package discovery

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/omalloc/kratos-agent/api/agent"
	"time"
)

func NewAgentService(logger log.Logger, dis registry.Discovery) (agent.AgentClient, error) {
	filter := FilterHang()
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos-agent"),
		grpc.WithDiscovery(dis),
		grpc.WithTimeout(time.Second*10),
		grpc.WithNodeFilter(filter),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
			tracing.Client(),
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	return agent.NewAgentClient(conn), nil
}
