package intf

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
	"net/http"
)

const (
	CtxTest_HANDLER_PATH = "/CtxTest"
)

type (
	CtxTestService interface {
		Exec(in *CtxTestIn) (*ykit.Result, error)
	}

	//input data
	CtxTestIn struct {
		S string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	CtxTestHandler struct {
		base ykit.RootTran
	}
)

func (r *CtxTestHandler) MakeLocalEndpoint(svc CtxTestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#######  Local Context ######  CtxTest ###########")
		ylog.Debug("jwt:", ctx.Value(ykit.JWT_TOKEN))
		spew.Dump(ctx)

		in := request.(*CtxTestIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *CtxTestHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	v, err := r.base.DecodeRequest(new(CtxTestIn), ctx, req)
	return v, err
}

//个人实现,参数不能修改
func (r *CtxTestHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *CtxTestHandler) HandlerLocal(service CtxTestService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	before := tran.ServerBefore(ykit.Jwt2ctx())

	opts := make([]tran.ServerOption, 0)
	opts = append(opts, before)
	opts = append(opts, options...)

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		opts...)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *CtxTestHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		CtxTest_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid)

}

func (r *CtxTestHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		CtxTest_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
