package ykit

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
	"net/http"
	"testing"
)

var f = new(RootTran).ProxyEndpointSDDefault(
	context.Background(),
	"user",
	"POST",
	"/hello",
	nil)

func Test_proxy_sd_call(t *testing.T) {
	w := g.NewWaitGroupN(100)
	for i := 0; i < 5000; i++ {
		w.Call(func() error {
			call(i, f)
			return nil
		})
	}
	ylog.Debug("time:", w.Wait())

}

func call(i int, f endpoint.Endpoint) {
	m := M{
		"S": A{"a", "b", "c", fmt.Sprint(i)}}

	y, err := f(context.Background(), m)
	if err != nil {
		ylog.Error("ykit_test.go->", err)
		return
	}
	log.Println("----------", "ok", "------------")
	spew.Dump(y)
}

func Test_GetUIDOfContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Uid", 123)
	uid := GetUIDOfContext(ctx)
	log.Println("----------", "uid: ", uid, "------------")

	ctx = context.WithValue(ctx, "Uid", "12345789")
	uid = GetUIDOfContext(ctx)
	log.Println("----------", "uid: ", uid, "------------")

}
func Test_GetUIDOfContext_jwt(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Jwt", "test/1234")
	uid := GetUIDOfContext(ctx)
	log.Println("----------", "uid: ", uid, "------------")
}

func Test_GetUIDOfReq(t *testing.T) {
	req := &http.Request{
		Header: map[string][]string{},
	}

	req.Header.Set(JWT_TOKEN, "test/123")
	uid := GetUIDOfReq(req)
	log.Println("----------", "uid : ", uid, "------------")
}
