package ykit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	tran "github.com/go-kit/kit/transport/http"

	"github.com/vhaoran/vchat/lib/yetcd"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ymid"
)

/*--auth: whr  date:2019-12-05--------------------------
 这是微服务是基类，所有实现微服务的接口都
 需要"继承"此类，用于快速开发及部署,
--------------------------------------- */
type (
	RootTran struct {
	}
)

//指定一个确定的数据类型，进行decode
func (r *RootTran) DecodeRequest(reqDataPtr interface{}, _ context.Context, req *http.Request) (interface{}, error) {
	//ylog.Debug("--------RootTran.go--->DeecodeRequest---")

	//for k, v := range req.URL.Query() {
	//	ylog.Debug("--------RootTran.go--->query param---", k, ":", v)
	//}

	if err := json.NewDecoder(req.Body).Decode(reqDataPtr); err != nil {
		ylog.Error("****** RootTran.go->DecodeRequest", err)
		if req.Method != "GET" {
			return nil, err
		}
	}

	//golog.Println("RootTran) DecodeRequest")
	//spew.Dump(reqDataPtr)
	ylog.Debug(fmt.Sprint("------url----", req.URL.Path, "------"))
	ylog.DebugDump("-----------传入参数(param)----", reqDataPtr)

	return reqDataPtr, nil
}

//针对特定传入类型的decode
func (r *RootTran) DecodeRequestDefault(ctx context.Context, req *http.Request) (interface{}, error) {
	//ylog.Debug("RootTran.go->DecodeRequestDefault")

	return r.DecodeRequest(new(RequestDefault), ctx, req)
}

func (r *RootTran) EncodeRequestBuffer(_ context.Context, req *http.Request, reqData interface{}) error {
	//ylog.Debug("RootTran.go->EncodeRequestBuffer")

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(reqData); err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func (r *RootTran) EncodeResponse(ctx context.Context, wr http.ResponseWriter, res interface{}) error {
	//ylog.Debug("RootTran.go->EncodeResponse")
	return json.NewEncoder(wr).Encode(res)
}

//实现decoder
func (r *RootTran) DecodeResponseDefault(_ context.Context, res *http.Response) (interface{}, error) {
	//ylog.Debug("--------RootTran.go--->DecodeResponseDefault---")

	var response Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		ylog.Error("RootTran.go->DecodeResponseDefault", err)
		//ylog.ErrorDump(res.Body)
		//ylog.Debug("----------------------------------")

		return nil, err
	}
	return response, nil
}

//实现decoder
func (r *RootTran) DecodeResponseString(_ context.Context, res *http.Response) (interface{}, error) {
	//ylog.Debug("--------RootTran.go--->DecodeResponseString---")
	buf := make([]byte, res.ContentLength*2)

	i, err := res.Body.Read(buf)
	if err != nil {
		err = errors.New(fmt.Sprint("错误：RootTran.go--DecodeResponseString:- ", err.Error()))
		ylog.Error(err.Error())
		return "", err
	}

	s := string(buf[0:i])
	return s, nil
}

//manual proxy
func (r *RootTran) ProxyEndPointOfInstance(
	instance,
	method,
	path string,
	decodeResponse func(_ context.Context, r *http.Response) (interface{}, error),
	ctx context.Context) endpoint.Endpoint {
	//instance := "localhost:9999"

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		ylog.Error("RootTran.go->ProxyEndPointOfInstance", err)
		panic(err)
	}
	u.Path = path

	//ylog.Debug("--------begin Visit:-------", instance, "->", path)

	return tran.NewClient(
		method,
		u,
		r.EncodeRequestBuffer,
		decodeResponse,
	).Endpoint()
}

//unit auto discovery
func (r *RootTran) ProxyEndpointSD(ctx context.Context,
	serviceTag, method, path string,
	decodeRequestFunc func(ctx context.Context, req *http.Request) (interface{}, error),
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error),
) endpoint.Endpoint {
	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 10 * 1000 * time.Millisecond

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		ylog.Error("--------RootTran.go--->etcdv3.NewClient-error--", err)
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		ylog.Error("--------RootTran.go--->etcdv3.NewInstancer-error--", err)
		return nil
	}

	//
	factory := r.FactorySD(ctx, method, path, decodeResponseFunc)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	ep := retry

	return ep
}

func (r *RootTran) ProxyEndpointSDDefault(ctx context.Context,
	serviceTag, method, path string,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) endpoint.Endpoint {

	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 10 * 1000 * time.Millisecond

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取etcd连接时失败，err:", err, " etcd config：", spew.Sdump(yetcd.XETCDConfig))
		ylog.Debug("RootTran.go->ProxyEndpointSDDefault,获取etcd连接时失败，err:", err)
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取实例时失败，err:", err)
		return nil
	}

	//
	factory := r.FactorySD(ctx,
		method,
		path,
		r.DecodeResponseDefault)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	ep := retry

	return ep
}

//unit auto discovery
func (r *RootTran) HandlerSD(ctx context.Context,
	serviceTag, method, path string,
	decodeRequestFunc func(ctx context.Context, req *http.Request) (interface{}, error),
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error),
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 10 * 1000 * time.Millisecond

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取etcd连接时失败，err:", err, " etcd config：", spew.Sdump(yetcd.XETCDConfig))
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取实例时失败，err:", err)
		return nil
	}

	//
	factory := r.FactorySD(ctx, method, path, decodeResponseFunc)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	ep := retry

	//------------bind middle ware-------------------
	ep = ymid.MidCommon(ep)
	for _, f := range mid {
		ep = f(ep)
	}

	//
	opts := append(options, tran.ServerBefore(Jwt2ctx()))
	opts = append(opts, tran.ServerBefore(QStr2ctx()))
	opts = append(opts, tran.ServerBefore(DebugHead()))

	return tran.NewServer(ep, decodeRequestFunc, r.EncodeResponse, opts...)
}

//unit auto discovery
//输入为map[string]interface{}
//输出为result的 handler
func (r *RootTran) HandlerSDDefault(ctx context.Context,
	serviceTag, method, path string,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 10 * 1000 * time.Millisecond

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取etcd连接时失败，err:", err, " etcd config：", spew.Sdump(yetcd.XETCDConfig))
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取实例时失败，err:", err)
		return nil
	}

	//
	factory := r.FactorySD(ctx,
		method,
		path,
		r.DecodeResponseDefault)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	ep := retry

	//bind middle ware
	ep = ymid.MidCommon(ep)
	for _, f := range mid {
		ep = f(ep)
	}

	opt := make([]tran.ServerOption, 0)
	opt = append(opt, ymid.ServerBeforeCallback)
	opt = append(opt, ymid.ServerBeforeCallback)

	opt = append(opt, options...)
	opt = append(opt, tran.ServerBefore(Jwt2ctx()))
	opt = append(opt, tran.ServerBefore(QStr2ctx()))

	return tran.NewServer(ep,
		r.DecodeRequestDefault,
		r.EncodeResponse,
		opt...)
}

//unit auto discovery
//输入为map[string]interface{}
//输出为result的 handler
func (r *RootTran) HandlerSDCommon(ctx context.Context,
	serviceTag, method, path string,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 10 * 1000 * time.Millisecond

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取etcd连接时失败，err:", err, " etcd config：", spew.Sdump(yetcd.XETCDConfig))
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		ylog.Error("RootTran.go->HandlerSD,获取实例时失败，err:", err)
		return nil
	}

	//
	factory := r.FactorySD(ctx,
		method,
		path,
		r.DecodeResponseDefault)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	ep := retry

	//bind middle ware
	ep = ymid.MidCommon(ep)
	for _, f := range mid {
		ep = f(ep)
	}

	opt := make([]tran.ServerOption, 0)
	opt = append(opt, ymid.ServerBeforeCallback)
	opt = append(opt, options...)
	opt = append(opt, tran.ServerBefore(Jwt2ctx()))
	opt = append(opt, tran.ServerBefore(QStr2ctx()))
	opt = append(opt, tran.ServerBefore(DebugHead()))
	//opt = append(opt, tran.ServerBefore(Head2Context()))

	return tran.NewServer(ep,
		r.DecodeRequestDefault,
		r.EncodeResponse,
		opt...)
}

// unit discovery
func (r *RootTran) FactorySD(
	ctx context.Context,
	method,
	path string,
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error),
	mid ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		//ylog.Debug("--------RootTran.go->factory->instance---", instance)

		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		targetURL, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		targetURL.Path = path

		//ylog.Debug("--------begin Visit:-------", instance, "->", path)

		enc := r.EncodeRequestBuffer
		dec := decodeResponseFunc

		opts := make([]tran.ClientOption, 0)
		opts = append(opts, tran.ClientBefore(Jwt2Req()))
		opts = append(opts, tran.ClientBefore(QStr2Req()))
		opts = append(opts, tran.ClientBefore(CommonHead()))
		opts = append(opts, tran.ClientBefore(DebugHead()))

		ep := tran.NewClient(method,
			targetURL,
			enc,
			dec,
			opts...).Endpoint()

		ep = ymid.MidCommon(ep)
		for _, f := range mid {
			ep = f(ep)
		}
		return ep, nil, nil
	}
}

func (r *RootTran) DecodeResponseMap(ctx context.Context, res *http.Response) (interface{}, error) {
	//ylog.Debug("RootTran.go->DecodeResponseMap")

	var response map[string]interface{} = make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		ylog.Error("RootTran.go->DecodeResponseMap", err)
		return nil, err
	}

	return response, nil
}
