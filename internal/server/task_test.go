package server_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/kratos/registry"
	"github.com/omalloc/contrib/protobuf"
	"github.com/omalloc/kratos-agent/api/agent"

	"github.com/omalloc/kratos-console/internal/discovery"
)

func TestTaskServer_flushed(t *testing.T) {
	etcdClient, cleanup1, err := registry.NewEtcd(&protobuf.Registry{
		Enabled:   true,
		Namespace: "/bedrock",
		Endpoints: []string{"172.18.80.234:2379"},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup1()

	dis := registry.NewDiscovery(etcdClient, &protobuf.Registry{
		Namespace: "/bedrock",
		Endpoints: []string{"172.18.80.234:2379"},
	})

	cli, err := discovery.NewAgentService(log.GetLogger(), dis)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	reply, err := cli.ListService(ctx, &agent.ListServiceRequest{})
	if err != nil {
		t.Fatal(err)
	}

	buf, err := json.MarshalIndent(reply, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(buf))
}
