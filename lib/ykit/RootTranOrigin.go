package ykit

import (
	"context"
	"io"
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

type (
	RootTranOrigin struct {
		RootTran
	}
)

func (r *RootTranOrigin) HandlerSDOriginDefault(ctx context.Context,
	serviceTag, method, path string,
	mid []endpoint.Middleware) *tran.Server {
	return r.HandlerSDOrigin(ctx,
		serviceTag,
		method,
		path,
		r.DecodeRequestDefault,
		r.DecodeResponseString,
		mid)
}

//unit auto discovery
func (r *RootTranOrigin) HandlerSDOrigin(ctx context.Context,
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

	//here
	factory := r.FactorySDOrigin(ctx, method, path, decodeResponseFunc)
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
	//here
	return tran.NewServer(ep, decodeRequestFunc, r.EncodeResponse, opts...)
}

func (r *RootTranOrigin) FactorySDOrigin(
	ctx context.Context,
	method,
	path string,
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error),
	mid ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
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
		//opts = append(opts, tran.ClientBefore(Jwt2Req()))
		opts = append(opts, tran.ClientBefore(OriginHead()))
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
