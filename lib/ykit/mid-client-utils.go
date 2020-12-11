package ykit

import (
	"context"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
	"net/http"
)

func Head2Context() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		ylog.Debug("########### enter Head2Context ######")
		for k, v := range req.Header {
			ylog.Debug("###### ", k, v, "####")
		}
		ylog.Debug("end Head2Context head Disp ")
		return context.WithValue(ctx, "Info", req.Header)
	}
}
func Context2Head() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		c := ctx
		m := ctx.Value("Info")
		//.(map[string][]string)
		log.Println("----------", "ctx content", "------------")
		ylog.DebugDump("context2Head ctx: ", ctx)
		log.Println("----------", "m-value", "------------")
		ylog.DebugDump("context2Head m:", m)

		if m == nil {
			return c
		}

		z, ok := m.(http.Header)
		if !ok {
			ylog.Debug("type is dis match")

			return c
		}

		for k, a := range z {
			ylog.Debug("******* set value:", k, ":", a)
			for _, v := range a {
				req.Header.Set(k, v)
			}
		}

		return c
	}
}
