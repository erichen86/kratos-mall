// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package v1

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

type PayHTTPServer interface {
	AppUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error)
	H5UnifiedOrder(context.Context, *UnifiedOrderReq) (*H5UnifiedOrderResp, error)
	JsUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error)
	OrderQuery(context.Context, *OrderQueryReq) (*OrderQueryResp, error)
}

func RegisterPayHTTPServer(s *http.Server, srv PayHTTPServer) {
	r := s.Route("/")
	r.POST("/api/pay/wx/appUnifiedOrder", _Pay_AppUnifiedOrder0_HTTP_Handler(srv))
	r.POST("/api/pay/wx/h5UnifiedOrder", _Pay_H5UnifiedOrder0_HTTP_Handler(srv))
	r.POST("/api/pay/wx/jsUnifiedOrder", _Pay_JsUnifiedOrder0_HTTP_Handler(srv))
	r.POST("/api/pay/wx/orderQuery", _Pay_OrderQuery0_HTTP_Handler(srv))
}

func _Pay_AppUnifiedOrder0_HTTP_Handler(srv PayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnifiedOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pay/AppUnifiedOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppUnifiedOrder(ctx, req.(*UnifiedOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnifiedOrderResp)
		return ctx.Result(200, reply)
	}
}

func _Pay_H5UnifiedOrder0_HTTP_Handler(srv PayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnifiedOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pay/H5UnifiedOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.H5UnifiedOrder(ctx, req.(*UnifiedOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*H5UnifiedOrderResp)
		return ctx.Result(200, reply)
	}
}

func _Pay_JsUnifiedOrder0_HTTP_Handler(srv PayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnifiedOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pay/JsUnifiedOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.JsUnifiedOrder(ctx, req.(*UnifiedOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnifiedOrderResp)
		return ctx.Result(200, reply)
	}
}

func _Pay_OrderQuery0_HTTP_Handler(srv PayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderQueryReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pay/OrderQuery")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderQuery(ctx, req.(*OrderQueryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderQueryResp)
		return ctx.Result(200, reply)
	}
}

type PayHTTPClient interface {
	AppUnifiedOrder(ctx context.Context, req *UnifiedOrderReq, opts ...http.CallOption) (rsp *UnifiedOrderResp, err error)
	H5UnifiedOrder(ctx context.Context, req *UnifiedOrderReq, opts ...http.CallOption) (rsp *H5UnifiedOrderResp, err error)
	JsUnifiedOrder(ctx context.Context, req *UnifiedOrderReq, opts ...http.CallOption) (rsp *UnifiedOrderResp, err error)
	OrderQuery(ctx context.Context, req *OrderQueryReq, opts ...http.CallOption) (rsp *OrderQueryResp, err error)
}

type PayHTTPClientImpl struct {
	cc *http.Client
}

func NewPayHTTPClient(client *http.Client) PayHTTPClient {
	return &PayHTTPClientImpl{client}
}

func (c *PayHTTPClientImpl) AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...http.CallOption) (*UnifiedOrderResp, error) {
	var out UnifiedOrderResp
	pattern := "/api/pay/wx/appUnifiedOrder"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pay/AppUnifiedOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PayHTTPClientImpl) H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...http.CallOption) (*H5UnifiedOrderResp, error) {
	var out H5UnifiedOrderResp
	pattern := "/api/pay/wx/h5UnifiedOrder"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pay/H5UnifiedOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PayHTTPClientImpl) JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...http.CallOption) (*UnifiedOrderResp, error) {
	var out UnifiedOrderResp
	pattern := "/api/pay/wx/jsUnifiedOrder"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pay/JsUnifiedOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PayHTTPClientImpl) OrderQuery(ctx context.Context, in *OrderQueryReq, opts ...http.CallOption) (*OrderQueryResp, error) {
	var out OrderQueryResp
	pattern := "/api/pay/wx/orderQuery"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pay/OrderQuery"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
