package ykit

import (
	"context"
	"net/http"
	"net/url"

	tran "github.com/go-kit/kit/transport/http"
)

const (
	Q_STR = "origin-q-str"
)

//query str to context
func QStr2ctx() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, Q_STR, req.URL.Query())
	}
}

//query string to context
func QStr2Req() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		m := ctx.Value(Q_STR)

		//
		v, ok := m.(url.Values)
		v.Encode()
		if ok && v != nil {
			req.URL.ForceQuery = true
			req.URL.RawQuery = v.Encode()
		}

		return ctx
	}
}
