// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

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

type PmsHTTPServer interface {
	BrandAdd(context.Context, *BrandAddReq) (*BrandAddResp, error)
	BrandDelete(context.Context, *BrandDeleteReq) (*BrandDeleteResp, error)
	BrandList(context.Context, *BrandListReq) (*BrandListResp, error)
	BrandUpdate(context.Context, *BrandUpdateReq) (*BrandUpdateResp, error)
	CommentAdd(context.Context, *CommentAddReq) (*CommentAddResp, error)
	CommentDelete(context.Context, *CommentDeleteReq) (*CommentDeleteResp, error)
	CommentList(context.Context, *CommentListReq) (*CommentListResp, error)
	CommentUpdate(context.Context, *CommentUpdateReq) (*CommentUpdateResp, error)
	FeightTemplateAdd(context.Context, *FeightTemplateAddReq) (*FeightTemplateAddResp, error)
	FeightTemplateDelete(context.Context, *FeightTemplateDeleteReq) (*FeightTemplateDeleteResp, error)
	FeightTemplateList(context.Context, *FeightTemplateListReq) (*FeightTemplateListResp, error)
	FeightTemplateUpdate(context.Context, *FeightTemplateUpdateReq) (*FeightTemplateUpdateResp, error)
	MemberPriceAdd(context.Context, *MemberPriceAddReq) (*MemberPriceAddResp, error)
	MemberPriceDelete(context.Context, *MemberPriceDeleteReq) (*MemberPriceDeleteResp, error)
	MemberPriceList(context.Context, *MemberPriceListReq) (*MemberPriceListResp, error)
	MemberPriceUpdate(context.Context, *MemberPriceUpdateReq) (*MemberPriceUpdateResp, error)
	ProductAdd(context.Context, *ProductAddReq) (*ProductAddResp, error)
	ProductCategoryAdd(context.Context, *ProductCategoryAddReq) (*ProductCategoryAddResp, error)
	ProductCategoryDelete(context.Context, *ProductCategoryDeleteReq) (*ProductCategoryDeleteResp, error)
	ProductCategoryList(context.Context, *ProductCategoryListReq) (*ProductCategoryListResp, error)
	ProductCategoryUpdate(context.Context, *ProductCategoryUpdateReq) (*ProductCategoryUpdateResp, error)
	ProductDelete(context.Context, *ProductDeleteReq) (*ProductDeleteResp, error)
	ProductList(context.Context, *ProductListReq) (*ProductListResp, error)
	ProductUpdate(context.Context, *ProductUpdateReq) (*ProductUpdateResp, error)
	SkuStockAdd(context.Context, *SkuStockAddReq) (*SkuStockAddResp, error)
	SkuStockDelete(context.Context, *SkuStockDeleteReq) (*SkuStockDeleteResp, error)
	SkuStockList(context.Context, *SkuStockListReq) (*SkuStockListResp, error)
	SkuStockUpdate(context.Context, *SkuStockUpdateReq) (*SkuStockUpdateResp, error)
}

func RegisterPmsHTTPServer(s *http.Server, srv PmsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/product/product/add", _Pms_ProductAdd0_HTTP_Handler(srv))
	r.POST("/api/product/product/list", _Pms_ProductList0_HTTP_Handler(srv))
	r.POST("/api/product/product/update", _Pms_ProductUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/product/delete", _Pms_ProductDelete0_HTTP_Handler(srv))
	r.POST("/api/product/brand/add", _Pms_BrandAdd0_HTTP_Handler(srv))
	r.POST("/api/product/brand/list", _Pms_BrandList0_HTTP_Handler(srv))
	r.POST("/api/product/brand/update", _Pms_BrandUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/brand/delete", _Pms_BrandDelete0_HTTP_Handler(srv))
	r.POST("/api/product/comment/add", _Pms_CommentAdd0_HTTP_Handler(srv))
	r.POST("/api/product/comment/list", _Pms_CommentList0_HTTP_Handler(srv))
	r.POST("/api/product/comment/update", _Pms_CommentUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/comment/delete", _Pms_CommentDelete0_HTTP_Handler(srv))
	r.POST("/api/product/feighttemplate/add", _Pms_FeightTemplateAdd0_HTTP_Handler(srv))
	r.POST("/api/product/feighttemplate/list", _Pms_FeightTemplateList0_HTTP_Handler(srv))
	r.POST("/api/product/feighttemplate/update", _Pms_FeightTemplateUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/feighttemplate/delete", _Pms_FeightTemplateDelete0_HTTP_Handler(srv))
	r.POST("/api/product/memberprice/add", _Pms_MemberPriceAdd0_HTTP_Handler(srv))
	r.POST("/api/product/memberprice/list", _Pms_MemberPriceList0_HTTP_Handler(srv))
	r.POST("/api/product/memberprice/update", _Pms_MemberPriceUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/memberprice/delete", _Pms_MemberPriceDelete0_HTTP_Handler(srv))
	r.POST("/api/product/category/add", _Pms_ProductCategoryAdd0_HTTP_Handler(srv))
	r.POST("/api/product/category/list", _Pms_ProductCategoryList0_HTTP_Handler(srv))
	r.POST("/api/product/category/update", _Pms_ProductCategoryUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/category/delete", _Pms_ProductCategoryDelete0_HTTP_Handler(srv))
	r.POST("/api/product/skustock/add", _Pms_SkuStockAdd0_HTTP_Handler(srv))
	r.POST("/api/product/skustock/list", _Pms_SkuStockList0_HTTP_Handler(srv))
	r.POST("/api/product/skustock/update", _Pms_SkuStockUpdate0_HTTP_Handler(srv))
	r.POST("/api/product/skustock/delete", _Pms_SkuStockDelete0_HTTP_Handler(srv))
}

func _Pms_ProductAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductAdd(ctx, req.(*ProductAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductList(ctx, req.(*ProductListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductUpdate(ctx, req.(*ProductUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductDelete(ctx, req.(*ProductDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_BrandAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BrandAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/BrandAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BrandAdd(ctx, req.(*BrandAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BrandAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_BrandList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BrandListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/BrandList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BrandList(ctx, req.(*BrandListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BrandListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_BrandUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BrandUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/BrandUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BrandUpdate(ctx, req.(*BrandUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BrandUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_BrandDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BrandDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/BrandDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BrandDelete(ctx, req.(*BrandDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BrandDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_CommentAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/CommentAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentAdd(ctx, req.(*CommentAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_CommentList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/CommentList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentList(ctx, req.(*CommentListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_CommentUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/CommentUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentUpdate(ctx, req.(*CommentUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_CommentDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/CommentDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentDelete(ctx, req.(*CommentDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_FeightTemplateAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeightTemplateAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/FeightTemplateAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeightTemplateAdd(ctx, req.(*FeightTemplateAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeightTemplateAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_FeightTemplateList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeightTemplateListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/FeightTemplateList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeightTemplateList(ctx, req.(*FeightTemplateListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeightTemplateListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_FeightTemplateUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeightTemplateUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/FeightTemplateUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeightTemplateUpdate(ctx, req.(*FeightTemplateUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeightTemplateUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_FeightTemplateDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeightTemplateDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/FeightTemplateDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeightTemplateDelete(ctx, req.(*FeightTemplateDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeightTemplateDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_MemberPriceAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberPriceAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/MemberPriceAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberPriceAdd(ctx, req.(*MemberPriceAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberPriceAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_MemberPriceList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberPriceListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/MemberPriceList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberPriceList(ctx, req.(*MemberPriceListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberPriceListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_MemberPriceUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberPriceUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/MemberPriceUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberPriceUpdate(ctx, req.(*MemberPriceUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberPriceUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_MemberPriceDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberPriceDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/MemberPriceDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberPriceDelete(ctx, req.(*MemberPriceDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberPriceDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductCategoryAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductCategoryAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductCategoryAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductCategoryAdd(ctx, req.(*ProductCategoryAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductCategoryAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductCategoryList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductCategoryListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductCategoryList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductCategoryList(ctx, req.(*ProductCategoryListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductCategoryListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductCategoryUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductCategoryUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductCategoryUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductCategoryUpdate(ctx, req.(*ProductCategoryUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductCategoryUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_ProductCategoryDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProductCategoryDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/ProductCategoryDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProductCategoryDelete(ctx, req.(*ProductCategoryDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProductCategoryDeleteResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_SkuStockAdd0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SkuStockAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/SkuStockAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SkuStockAdd(ctx, req.(*SkuStockAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SkuStockAddResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_SkuStockList0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SkuStockListReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/SkuStockList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SkuStockList(ctx, req.(*SkuStockListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SkuStockListResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_SkuStockUpdate0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SkuStockUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/SkuStockUpdate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SkuStockUpdate(ctx, req.(*SkuStockUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SkuStockUpdateResp)
		return ctx.Result(200, reply)
	}
}

func _Pms_SkuStockDelete0_HTTP_Handler(srv PmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SkuStockDeleteReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/front.admin.v1.Pms/SkuStockDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SkuStockDelete(ctx, req.(*SkuStockDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SkuStockDeleteResp)
		return ctx.Result(200, reply)
	}
}

type PmsHTTPClient interface {
	BrandAdd(ctx context.Context, req *BrandAddReq, opts ...http.CallOption) (rsp *BrandAddResp, err error)
	BrandDelete(ctx context.Context, req *BrandDeleteReq, opts ...http.CallOption) (rsp *BrandDeleteResp, err error)
	BrandList(ctx context.Context, req *BrandListReq, opts ...http.CallOption) (rsp *BrandListResp, err error)
	BrandUpdate(ctx context.Context, req *BrandUpdateReq, opts ...http.CallOption) (rsp *BrandUpdateResp, err error)
	CommentAdd(ctx context.Context, req *CommentAddReq, opts ...http.CallOption) (rsp *CommentAddResp, err error)
	CommentDelete(ctx context.Context, req *CommentDeleteReq, opts ...http.CallOption) (rsp *CommentDeleteResp, err error)
	CommentList(ctx context.Context, req *CommentListReq, opts ...http.CallOption) (rsp *CommentListResp, err error)
	CommentUpdate(ctx context.Context, req *CommentUpdateReq, opts ...http.CallOption) (rsp *CommentUpdateResp, err error)
	FeightTemplateAdd(ctx context.Context, req *FeightTemplateAddReq, opts ...http.CallOption) (rsp *FeightTemplateAddResp, err error)
	FeightTemplateDelete(ctx context.Context, req *FeightTemplateDeleteReq, opts ...http.CallOption) (rsp *FeightTemplateDeleteResp, err error)
	FeightTemplateList(ctx context.Context, req *FeightTemplateListReq, opts ...http.CallOption) (rsp *FeightTemplateListResp, err error)
	FeightTemplateUpdate(ctx context.Context, req *FeightTemplateUpdateReq, opts ...http.CallOption) (rsp *FeightTemplateUpdateResp, err error)
	MemberPriceAdd(ctx context.Context, req *MemberPriceAddReq, opts ...http.CallOption) (rsp *MemberPriceAddResp, err error)
	MemberPriceDelete(ctx context.Context, req *MemberPriceDeleteReq, opts ...http.CallOption) (rsp *MemberPriceDeleteResp, err error)
	MemberPriceList(ctx context.Context, req *MemberPriceListReq, opts ...http.CallOption) (rsp *MemberPriceListResp, err error)
	MemberPriceUpdate(ctx context.Context, req *MemberPriceUpdateReq, opts ...http.CallOption) (rsp *MemberPriceUpdateResp, err error)
	ProductAdd(ctx context.Context, req *ProductAddReq, opts ...http.CallOption) (rsp *ProductAddResp, err error)
	ProductCategoryAdd(ctx context.Context, req *ProductCategoryAddReq, opts ...http.CallOption) (rsp *ProductCategoryAddResp, err error)
	ProductCategoryDelete(ctx context.Context, req *ProductCategoryDeleteReq, opts ...http.CallOption) (rsp *ProductCategoryDeleteResp, err error)
	ProductCategoryList(ctx context.Context, req *ProductCategoryListReq, opts ...http.CallOption) (rsp *ProductCategoryListResp, err error)
	ProductCategoryUpdate(ctx context.Context, req *ProductCategoryUpdateReq, opts ...http.CallOption) (rsp *ProductCategoryUpdateResp, err error)
	ProductDelete(ctx context.Context, req *ProductDeleteReq, opts ...http.CallOption) (rsp *ProductDeleteResp, err error)
	ProductList(ctx context.Context, req *ProductListReq, opts ...http.CallOption) (rsp *ProductListResp, err error)
	ProductUpdate(ctx context.Context, req *ProductUpdateReq, opts ...http.CallOption) (rsp *ProductUpdateResp, err error)
	SkuStockAdd(ctx context.Context, req *SkuStockAddReq, opts ...http.CallOption) (rsp *SkuStockAddResp, err error)
	SkuStockDelete(ctx context.Context, req *SkuStockDeleteReq, opts ...http.CallOption) (rsp *SkuStockDeleteResp, err error)
	SkuStockList(ctx context.Context, req *SkuStockListReq, opts ...http.CallOption) (rsp *SkuStockListResp, err error)
	SkuStockUpdate(ctx context.Context, req *SkuStockUpdateReq, opts ...http.CallOption) (rsp *SkuStockUpdateResp, err error)
}

type PmsHTTPClientImpl struct {
	cc *http.Client
}

func NewPmsHTTPClient(client *http.Client) PmsHTTPClient {
	return &PmsHTTPClientImpl{client}
}

func (c *PmsHTTPClientImpl) BrandAdd(ctx context.Context, in *BrandAddReq, opts ...http.CallOption) (*BrandAddResp, error) {
	var out BrandAddResp
	pattern := "/api/product/brand/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/BrandAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) BrandDelete(ctx context.Context, in *BrandDeleteReq, opts ...http.CallOption) (*BrandDeleteResp, error) {
	var out BrandDeleteResp
	pattern := "/api/product/brand/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/BrandDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) BrandList(ctx context.Context, in *BrandListReq, opts ...http.CallOption) (*BrandListResp, error) {
	var out BrandListResp
	pattern := "/api/product/brand/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/BrandList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) BrandUpdate(ctx context.Context, in *BrandUpdateReq, opts ...http.CallOption) (*BrandUpdateResp, error) {
	var out BrandUpdateResp
	pattern := "/api/product/brand/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/BrandUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) CommentAdd(ctx context.Context, in *CommentAddReq, opts ...http.CallOption) (*CommentAddResp, error) {
	var out CommentAddResp
	pattern := "/api/product/comment/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/CommentAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) CommentDelete(ctx context.Context, in *CommentDeleteReq, opts ...http.CallOption) (*CommentDeleteResp, error) {
	var out CommentDeleteResp
	pattern := "/api/product/comment/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/CommentDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) CommentList(ctx context.Context, in *CommentListReq, opts ...http.CallOption) (*CommentListResp, error) {
	var out CommentListResp
	pattern := "/api/product/comment/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/CommentList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) CommentUpdate(ctx context.Context, in *CommentUpdateReq, opts ...http.CallOption) (*CommentUpdateResp, error) {
	var out CommentUpdateResp
	pattern := "/api/product/comment/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/CommentUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) FeightTemplateAdd(ctx context.Context, in *FeightTemplateAddReq, opts ...http.CallOption) (*FeightTemplateAddResp, error) {
	var out FeightTemplateAddResp
	pattern := "/api/product/feighttemplate/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/FeightTemplateAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) FeightTemplateDelete(ctx context.Context, in *FeightTemplateDeleteReq, opts ...http.CallOption) (*FeightTemplateDeleteResp, error) {
	var out FeightTemplateDeleteResp
	pattern := "/api/product/feighttemplate/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/FeightTemplateDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) FeightTemplateList(ctx context.Context, in *FeightTemplateListReq, opts ...http.CallOption) (*FeightTemplateListResp, error) {
	var out FeightTemplateListResp
	pattern := "/api/product/feighttemplate/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/FeightTemplateList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) FeightTemplateUpdate(ctx context.Context, in *FeightTemplateUpdateReq, opts ...http.CallOption) (*FeightTemplateUpdateResp, error) {
	var out FeightTemplateUpdateResp
	pattern := "/api/product/feighttemplate/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/FeightTemplateUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) MemberPriceAdd(ctx context.Context, in *MemberPriceAddReq, opts ...http.CallOption) (*MemberPriceAddResp, error) {
	var out MemberPriceAddResp
	pattern := "/api/product/memberprice/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/MemberPriceAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) MemberPriceDelete(ctx context.Context, in *MemberPriceDeleteReq, opts ...http.CallOption) (*MemberPriceDeleteResp, error) {
	var out MemberPriceDeleteResp
	pattern := "/api/product/memberprice/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/MemberPriceDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) MemberPriceList(ctx context.Context, in *MemberPriceListReq, opts ...http.CallOption) (*MemberPriceListResp, error) {
	var out MemberPriceListResp
	pattern := "/api/product/memberprice/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/MemberPriceList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) MemberPriceUpdate(ctx context.Context, in *MemberPriceUpdateReq, opts ...http.CallOption) (*MemberPriceUpdateResp, error) {
	var out MemberPriceUpdateResp
	pattern := "/api/product/memberprice/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/MemberPriceUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductAdd(ctx context.Context, in *ProductAddReq, opts ...http.CallOption) (*ProductAddResp, error) {
	var out ProductAddResp
	pattern := "/api/product/product/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductCategoryAdd(ctx context.Context, in *ProductCategoryAddReq, opts ...http.CallOption) (*ProductCategoryAddResp, error) {
	var out ProductCategoryAddResp
	pattern := "/api/product/category/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductCategoryAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductCategoryDelete(ctx context.Context, in *ProductCategoryDeleteReq, opts ...http.CallOption) (*ProductCategoryDeleteResp, error) {
	var out ProductCategoryDeleteResp
	pattern := "/api/product/category/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductCategoryDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductCategoryList(ctx context.Context, in *ProductCategoryListReq, opts ...http.CallOption) (*ProductCategoryListResp, error) {
	var out ProductCategoryListResp
	pattern := "/api/product/category/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductCategoryList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductCategoryUpdate(ctx context.Context, in *ProductCategoryUpdateReq, opts ...http.CallOption) (*ProductCategoryUpdateResp, error) {
	var out ProductCategoryUpdateResp
	pattern := "/api/product/category/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductCategoryUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...http.CallOption) (*ProductDeleteResp, error) {
	var out ProductDeleteResp
	pattern := "/api/product/product/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductList(ctx context.Context, in *ProductListReq, opts ...http.CallOption) (*ProductListResp, error) {
	var out ProductListResp
	pattern := "/api/product/product/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...http.CallOption) (*ProductUpdateResp, error) {
	var out ProductUpdateResp
	pattern := "/api/product/product/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/ProductUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) SkuStockAdd(ctx context.Context, in *SkuStockAddReq, opts ...http.CallOption) (*SkuStockAddResp, error) {
	var out SkuStockAddResp
	pattern := "/api/product/skustock/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/SkuStockAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) SkuStockDelete(ctx context.Context, in *SkuStockDeleteReq, opts ...http.CallOption) (*SkuStockDeleteResp, error) {
	var out SkuStockDeleteResp
	pattern := "/api/product/skustock/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/SkuStockDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) SkuStockList(ctx context.Context, in *SkuStockListReq, opts ...http.CallOption) (*SkuStockListResp, error) {
	var out SkuStockListResp
	pattern := "/api/product/skustock/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/SkuStockList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PmsHTTPClientImpl) SkuStockUpdate(ctx context.Context, in *SkuStockUpdateReq, opts ...http.CallOption) (*SkuStockUpdateResp, error) {
	var out SkuStockUpdateResp
	pattern := "/api/product/skustock/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/front.admin.v1.Pms/SkuStockUpdate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
