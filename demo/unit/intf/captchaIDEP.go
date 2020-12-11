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
	"net/http"
)

const (
	CaptchaID_H_PATH = "/CaptchaID"
)

type (
	CaptchaIDService interface {
		Exec(ctx context.Context, in *CaptchaIDIn) (*ykit.Result, error)
	}

	//input data
	CaptchaIDIn struct {
		S string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	CaptchaIDH struct {
		base ykit.RootTran
	}
)

func (r *CaptchaIDH) MakeLocalEndpoint(svc CaptchaIDService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  CaptchaID ###########")
		spew.Dump(ctx)

		in := request.(*CaptchaIDIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *CaptchaIDH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return new(CaptchaIDIn), nil
}

//个人实现,参数不能修改
func (r *CaptchaIDH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *CaptchaIDH) HandlerLocal(service CaptchaIDService,
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
func (r *CaptchaIDH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		CaptchaID_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *CaptchaIDH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		CaptchaID_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
