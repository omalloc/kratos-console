// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v5.27.1
// source: console/discovery/disocvery.proto

package discovery

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDiscoveryKVGetValue = "/api.console.discovery.Discovery/KVGetValue"
const OperationDiscoveryKVListClusters = "/api.console.discovery.Discovery/KVListClusters"
const OperationDiscoveryKVListKeys = "/api.console.discovery.Discovery/KVListKeys"
const OperationDiscoveryKVUpdateHang = "/api.console.discovery.Discovery/KVUpdateHang"
const OperationDiscoveryOnlineServices = "/api.console.discovery.Discovery/OnlineServices"

type DiscoveryHTTPServer interface {
	KVGetValue(context.Context, *KVGetValueRequest) (*KVGetValueReply, error)
	KVListClusters(context.Context, *KVListClustersRequest) (*KVListClustersReply, error)
	KVListKeys(context.Context, *KVListKeysRequest) (*KVListKeysReply, error)
	KVUpdateHang(context.Context, *KVUpdateHangRequest) (*KVUpdateHangReply, error)
	OnlineServices(context.Context, *OnlineServiceRequest) (*OnlineServiceReply, error)
}

func RegisterDiscoveryHTTPServer(s *http.Server, srv DiscoveryHTTPServer) {
	r := s.Route("/")
	r.GET("/api/console/discovery/services/-/online", _Discovery_OnlineServices0_HTTP_Handler(srv))
	r.GET("/api/console/kv/clusters", _Discovery_KVListClusters0_HTTP_Handler(srv))
	r.GET("/api/console/kv/keys", _Discovery_KVListKeys0_HTTP_Handler(srv))
	r.GET("/api/console/kv/-/value", _Discovery_KVGetValue0_HTTP_Handler(srv))
	r.PUT("/api/console/kv/{id}/hang", _Discovery_KVUpdateHang0_HTTP_Handler(srv))
}

func _Discovery_OnlineServices0_HTTP_Handler(srv DiscoveryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OnlineServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDiscoveryOnlineServices)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OnlineServices(ctx, req.(*OnlineServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OnlineServiceReply)
		return ctx.Result(200, reply)
	}
}

func _Discovery_KVListClusters0_HTTP_Handler(srv DiscoveryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KVListClustersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDiscoveryKVListClusters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KVListClusters(ctx, req.(*KVListClustersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KVListClustersReply)
		return ctx.Result(200, reply)
	}
}

func _Discovery_KVListKeys0_HTTP_Handler(srv DiscoveryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KVListKeysRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDiscoveryKVListKeys)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KVListKeys(ctx, req.(*KVListKeysRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KVListKeysReply)
		return ctx.Result(200, reply)
	}
}

func _Discovery_KVGetValue0_HTTP_Handler(srv DiscoveryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KVGetValueRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDiscoveryKVGetValue)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KVGetValue(ctx, req.(*KVGetValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KVGetValueReply)
		return ctx.Result(200, reply)
	}
}

func _Discovery_KVUpdateHang0_HTTP_Handler(srv DiscoveryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KVUpdateHangRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDiscoveryKVUpdateHang)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KVUpdateHang(ctx, req.(*KVUpdateHangRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KVUpdateHangReply)
		return ctx.Result(200, reply)
	}
}

type DiscoveryHTTPClient interface {
	KVGetValue(ctx context.Context, req *KVGetValueRequest, opts ...http.CallOption) (rsp *KVGetValueReply, err error)
	KVListClusters(ctx context.Context, req *KVListClustersRequest, opts ...http.CallOption) (rsp *KVListClustersReply, err error)
	KVListKeys(ctx context.Context, req *KVListKeysRequest, opts ...http.CallOption) (rsp *KVListKeysReply, err error)
	KVUpdateHang(ctx context.Context, req *KVUpdateHangRequest, opts ...http.CallOption) (rsp *KVUpdateHangReply, err error)
	OnlineServices(ctx context.Context, req *OnlineServiceRequest, opts ...http.CallOption) (rsp *OnlineServiceReply, err error)
}

type DiscoveryHTTPClientImpl struct {
	cc *http.Client
}

func NewDiscoveryHTTPClient(client *http.Client) DiscoveryHTTPClient {
	return &DiscoveryHTTPClientImpl{client}
}

func (c *DiscoveryHTTPClientImpl) KVGetValue(ctx context.Context, in *KVGetValueRequest, opts ...http.CallOption) (*KVGetValueReply, error) {
	var out KVGetValueReply
	pattern := "/api/console/kv/-/value"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDiscoveryKVGetValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DiscoveryHTTPClientImpl) KVListClusters(ctx context.Context, in *KVListClustersRequest, opts ...http.CallOption) (*KVListClustersReply, error) {
	var out KVListClustersReply
	pattern := "/api/console/kv/clusters"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDiscoveryKVListClusters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DiscoveryHTTPClientImpl) KVListKeys(ctx context.Context, in *KVListKeysRequest, opts ...http.CallOption) (*KVListKeysReply, error) {
	var out KVListKeysReply
	pattern := "/api/console/kv/keys"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDiscoveryKVListKeys))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DiscoveryHTTPClientImpl) KVUpdateHang(ctx context.Context, in *KVUpdateHangRequest, opts ...http.CallOption) (*KVUpdateHangReply, error) {
	var out KVUpdateHangReply
	pattern := "/api/console/kv/{id}/hang"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDiscoveryKVUpdateHang))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DiscoveryHTTPClientImpl) OnlineServices(ctx context.Context, in *OnlineServiceRequest, opts ...http.CallOption) (*OnlineServiceReply, error) {
	var out OnlineServiceReply
	pattern := "/api/console/discovery/services/-/online"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDiscoveryOnlineServices))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
