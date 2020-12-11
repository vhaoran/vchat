package intf

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vhaoran/vchat/common/ypage"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	PB_H_PATH = "/PB"
)

/*
curl -X POST  \
-H 'Content-Type:application/json' \
-H 'Jwt:test/1' \
-d '{"page_no":1,"rows_per_page":2}' \
127.0.0.1:9999/api/pb \



*/

type (
	PBService interface {
		Exec(ctx context.Context, in *PBIn) (*ykit.Result, error)
	}

	//input data
	PBIn struct {
		ypage.PageBean
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	PBH struct {
		base ykit.RootTran
	}
)

func (r *PBH) MakeLocalEndpoint(svc PBService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  PB ###########")
		spew.Dump(ctx)

		in := request.(*PBIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *PBH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	log.Println("----------", "------------")
	p := make([]byte, 1024*4)
	c, err := req.Body.Read(p)

	log.Println("-err--", c, "---", err, "------------")
	log.Println(string(p))

	return r.base.DecodeRequest(new(PBIn), ctx, req)
}

//个人实现,参数不能修改
func (r *PBH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *PBH) HandlerLocal(service PBService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	opts := make([]tran.ServerOption, 0)
	opts = append(opts, tran.ServerBefore(ykit.Jwt2ctx()))
	opts = append(opts, tran.ServerBefore(ykit.DebugHead()))
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
func (r *PBH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		PB_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *PBH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		PB_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
