// Code generated by protoc-gen-go-http-kitx. DO NOT EDIT.
// versions:
// - protoc-gen-go-http-kitx v0.0.1
// - protoc             v3.17.3
// source: greeter.proto

package v1

import (
	context "context"
	http "github.com/sado0823/go-kitx/transport/http"
	binding "github.com/sado0823/go-kitx/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kitx package it is being compiled against.
var _ = new(context.Context)
var _ http.Transporter = (*http.Transport)(nil)
var _ = binding.EncodeURL

const OperationGreeterAddBook = "/helloworld.v1.Greeter/AddBook"
const OperationGreeterSayHello = "/helloworld.v1.Greeter/SayHello"
const OperationGreeterShopList = "/helloworld.v1.Greeter/ShopList"

type GreeterHTTPServer interface {
	AddBook(context.Context, *AddBookRequest) (*AddBookReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	ShopList(context.Context, *ShopListRequest) (*ShopListReply, error)
}

func RegisterGreeterHTTPServer(r *http.Router, srv GreeterHTTPServer) {
	r.GET("/helloworld/{name}", _Greeter_SayHello0_HTTP_Handler(srv))
	r.POST("/book/add", _Greeter_AddBook0_HTTP_Handler(srv))
	r.GET("/shop/list", _Greeter_ShopList0_HTTP_Handler(srv))
}

func _Greeter_SayHello0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_AddBook0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddBookRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterAddBook)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddBook(ctx, req.(*AddBookRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddBookReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_ShopList0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShopListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterShopList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShopList(ctx, req.(*ShopListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ShopListReply)
		return ctx.Result(200, reply)
	}
}

type GreeterHTTPClient interface {
	AddBook(ctx context.Context, req *AddBookRequest, opts ...http.CallOption) (rsp *AddBookReply, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	ShopList(ctx context.Context, req *ShopListRequest, opts ...http.CallOption) (rsp *ShopListReply, err error)
}

type GreeterHTTPClientImpl struct {
	cc *http.Client
}

func NewGreeterHTTPClient(client *http.Client) GreeterHTTPClient {
	return &GreeterHTTPClientImpl{client}
}

func (c *GreeterHTTPClientImpl) AddBook(ctx context.Context, in *AddBookRequest, opts ...http.CallOption) (*AddBookReply, error) {
	var out AddBookReply
	pattern := "/book/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterAddBook))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) ShopList(ctx context.Context, in *ShopListRequest, opts ...http.CallOption) (*ShopListReply, error) {
	var out ShopListReply
	pattern := "/shop/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterShopList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
