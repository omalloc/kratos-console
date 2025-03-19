package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	adminpb "github.com/omalloc/kratos-console/api/console/administration"
	discoverypb "github.com/omalloc/kratos-console/api/console/discovery"
	resourcepb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/conf"
	"github.com/omalloc/kratos-console/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger,
	// resouces
	zone *service.ZoneService,
	node *service.NodeService,
	app *service.AppService,
	namespace *service.NamespaceService,

	// admin
	user *service.UserService,
	role *service.RoleService,
	permission *service.PermissionService,
	// discovery
	discovery *service.DiscoveryService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	resourcepb.RegisterZoneServer(srv, zone)
	resourcepb.RegisterNodeServer(srv, node)
	resourcepb.RegisterAppServer(srv, app)
	resourcepb.RegisterNamespaceServer(srv, namespace)
	discoverypb.RegisterDiscoveryServer(srv, discovery)

	adminpb.RegisterUserServer(srv, user)
	adminpb.RegisterRoleServer(srv, role)
	adminpb.RegisterPermissionServer(srv, permission)
	return srv
}
