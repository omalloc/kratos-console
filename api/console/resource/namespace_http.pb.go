// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: console/resource/namespace.proto

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

const OperationNamespaceCreate = "/api.console.resource.Namespace/Create"
const OperationNamespaceDelete = "/api.console.resource.Namespace/Delete"
const OperationNamespaceGet = "/api.console.resource.Namespace/Get"
const OperationNamespaceList = "/api.console.resource.Namespace/List"
const OperationNamespaceSimpleList = "/api.console.resource.Namespace/SimpleList"
const OperationNamespaceUpdate = "/api.console.resource.Namespace/Update"

type NamespaceHTTPServer interface {
	Create(context.Context, *NamespaceCreateRequest) (*NamespaceCreateReply, error)
	Delete(context.Context, *NamespaceDeleteRequest) (*NamespaceDeleteReply, error)
	Get(context.Context, *NamespaceGetRequest) (*NamespaceGetReply, error)
	List(context.Context, *NamespaceListRequest) (*NamespaceListReply, error)
	SimpleList(context.Context, *NamespaceSimpleListRequest) (*NamespaceSimpleListReply, error)
	Update(context.Context, *NamespaceUpdateRequest) (*NamespaceUpdateReply, error)
}

func RegisterNamespaceHTTPServer(s *http.Server, srv NamespaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/console/namespace", _Namespace_List1_HTTP_Handler(srv))
	r.GET("/api/console/namespace/-/options", _Namespace_SimpleList1_HTTP_Handler(srv))
	r.GET("/api/console/namespace/{id}", _Namespace_Get1_HTTP_Handler(srv))
	r.POST("/api/console/namespace", _Namespace_Create1_HTTP_Handler(srv))
	r.PUT("/api/console/namespace/{id}", _Namespace_Update1_HTTP_Handler(srv))
	r.DELETE("/api/console/namespace/{id}", _Namespace_Delete1_HTTP_Handler(srv))
}

func _Namespace_List1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*NamespaceListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceListReply)
		return ctx.Result(200, reply)
	}
}

func _Namespace_SimpleList1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceSimpleListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceSimpleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SimpleList(ctx, req.(*NamespaceSimpleListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceSimpleListReply)
		return ctx.Result(200, reply)
	}
}

func _Namespace_Get1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceGetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*NamespaceGetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceGetReply)
		return ctx.Result(200, reply)
	}
}

func _Namespace_Create1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*NamespaceCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceCreateReply)
		return ctx.Result(200, reply)
	}
}

func _Namespace_Update1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*NamespaceUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _Namespace_Delete1_HTTP_Handler(srv NamespaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNamespaceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*NamespaceDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceDeleteReply)
		return ctx.Result(200, reply)
	}
}

type NamespaceHTTPClient interface {
	Create(ctx context.Context, req *NamespaceCreateRequest, opts ...http.CallOption) (rsp *NamespaceCreateReply, err error)
	Delete(ctx context.Context, req *NamespaceDeleteRequest, opts ...http.CallOption) (rsp *NamespaceDeleteReply, err error)
	Get(ctx context.Context, req *NamespaceGetRequest, opts ...http.CallOption) (rsp *NamespaceGetReply, err error)
	List(ctx context.Context, req *NamespaceListRequest, opts ...http.CallOption) (rsp *NamespaceListReply, err error)
	SimpleList(ctx context.Context, req *NamespaceSimpleListRequest, opts ...http.CallOption) (rsp *NamespaceSimpleListReply, err error)
	Update(ctx context.Context, req *NamespaceUpdateRequest, opts ...http.CallOption) (rsp *NamespaceUpdateReply, err error)
}

type NamespaceHTTPClientImpl struct {
	cc *http.Client
}

func NewNamespaceHTTPClient(client *http.Client) NamespaceHTTPClient {
	return &NamespaceHTTPClientImpl{client}
}

func (c *NamespaceHTTPClientImpl) Create(ctx context.Context, in *NamespaceCreateRequest, opts ...http.CallOption) (*NamespaceCreateReply, error) {
	var out NamespaceCreateReply
	pattern := "/api/console/namespace"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNamespaceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NamespaceHTTPClientImpl) Delete(ctx context.Context, in *NamespaceDeleteRequest, opts ...http.CallOption) (*NamespaceDeleteReply, error) {
	var out NamespaceDeleteReply
	pattern := "/api/console/namespace/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNamespaceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NamespaceHTTPClientImpl) Get(ctx context.Context, in *NamespaceGetRequest, opts ...http.CallOption) (*NamespaceGetReply, error) {
	var out NamespaceGetReply
	pattern := "/api/console/namespace/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNamespaceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NamespaceHTTPClientImpl) List(ctx context.Context, in *NamespaceListRequest, opts ...http.CallOption) (*NamespaceListReply, error) {
	var out NamespaceListReply
	pattern := "/api/console/namespace"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNamespaceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NamespaceHTTPClientImpl) SimpleList(ctx context.Context, in *NamespaceSimpleListRequest, opts ...http.CallOption) (*NamespaceSimpleListReply, error) {
	var out NamespaceSimpleListReply
	pattern := "/api/console/namespace/-/options"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNamespaceSimpleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NamespaceHTTPClientImpl) Update(ctx context.Context, in *NamespaceUpdateRequest, opts ...http.CallOption) (*NamespaceUpdateReply, error) {
	var out NamespaceUpdateReply
	pattern := "/api/console/namespace/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNamespaceUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
