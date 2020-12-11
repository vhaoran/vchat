package ykit

import (
	"context"
	"net/http"

	tran "github.com/go-kit/kit/transport/http"

	"github.com/vhaoran/vchat/lib/ylog"
)

//不必yyoc原来的head,无条件转发
func OriginHead() tran.RequestFunc {

	return func(ctx context.Context, req *http.Request) context.Context {
		ylog.Debug("--------visit url--->---", req.URL.String())

		c := context.WithValue(ctx, "origin-header", req.Header)
		return c
	}
}
