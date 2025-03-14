package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	discoverypb "github.com/omalloc/kratos-console/api/console/discovery"
	resourcepb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/conf"
	"github.com/omalloc/kratos-console/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	// resources
	zone *service.ZoneService,
	node *service.NodeService,
	app *service.AppService,
	namespace *service.NamespaceService,
	// discovery
	discovery *service.DiscoveryService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/q/", openapiv2.NewHandler())

	resourcepb.RegisterZoneHTTPServer(srv, zone)
	resourcepb.RegisterNodeHTTPServer(srv, node)
	resourcepb.RegisterAppHTTPServer(srv, app)
	resourcepb.RegisterNamespaceHTTPServer(srv, namespace)
	discoverypb.RegisterDiscoveryHTTPServer(srv, discovery)
	return srv
}
