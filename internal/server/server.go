package server

import (
	"context"
	"github.com/google/wire"
	"github.com/omalloc/contrib/kratos/health"
	"github.com/omalloc/contrib/kratos/registry"
	"github.com/omalloc/contrib/protobuf"
	"github.com/omalloc/kratos-console/internal/conf"
	"github.com/omalloc/kratos-console/internal/data"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	NewRegistryConfig,
	NewChecker,

	registry.NewEtcd,
	registry.NewRegistrar,
	registry.NewDiscovery,

	health.NewServer,
)

func NewRegistryConfig(bc *conf.Bootstrap) *protobuf.Registry {
	return bc.Registry
}

func NewTracingConfig(bc *conf.Bootstrap) *protobuf.Tracing {
	return bc.Tracing
}

func NewChecker(c1 *data.Data, cli *clientv3.Client) []health.Checker {
	etcdChecker := &Etcd{cli}
	return []health.Checker{c1, etcdChecker}
}

type Etcd struct {
	*clientv3.Client
}

func (r *Etcd) Check(ctx context.Context) error {
	_, err := r.MemberList(ctx)
	return err
}
