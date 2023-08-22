// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.4
// source: console/resource/app.proto

package resource

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

const OperationAppList = "/api.console.resource.App/List"

type AppHTTPServer interface {
	List(context.Context, *AppListRequest) (*AppListReply, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/api/console/app", _App_List0_HTTP_Handler(srv))
}

func _App_List0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*AppListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppListReply)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	List(ctx context.Context, req *AppListRequest, opts ...http.CallOption) (rsp *AppListReply, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) List(ctx context.Context, in *AppListRequest, opts ...http.CallOption) (*AppListReply, error) {
	var out AppListReply
	pattern := "/api/console/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
