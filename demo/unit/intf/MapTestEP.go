package intf

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	MapTest_H_PATH = "/MapTest"
)

type (
	MapTestService interface {
		Exec(ctx context.Context, in *MapTestIn) (*MapTestOut, error)
	}

	//input data
	MapTestIn struct {
		S string `json:"s"`
	}

	//output data
	MapTestOut struct {
		A int    `json:"a,omitempty"`
		B string `json:"b,omitempty"`
		C int64  `json:"c,omitempty"`

		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	// handler implements
	MapTestH struct {
		base ykit.RootTran
	}
)

func (r *MapTestH) MakeLocalEndpoint(svc MapTestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  MapTest ###########")
		spew.Dump(ctx)

		in := request.(*MapTestIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *MapTestH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(MapTestIn), ctx, req)
}

//个人实现,参数不能修改
func (r *MapTestH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *MapTestOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *MapTestH) HandlerLocal(service MapTestService,
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
func (r *MapTestH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		MapTest_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *MapTestH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		MapTest_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_MapTest sync.Once
var local_MapTest_EP endpoint.Endpoint

func (r *MapTestH) Call(in MapTestIn) (*MapTestOut, error) {
	once_MapTest.Do(func() {
		local_MapTest_EP = new(MapTestH).ProxySD()
	})
	//
	ep := local_MapTest_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*MapTestOut), nil
}
