package ymid

import (
	"context"
	"net/http"

	tran "github.com/go-kit/kit/transport/http"

	"github.com/vhaoran/vchat/lib/ylog"
)

var ServerBeforeCallback = tran.ServerBefore(beforeCallback)

//黑夜的内置before,userd in <ServerOptions>
func beforeCallback(ctx context.Context, req *http.Request) context.Context {
	ylog.Info("-#############-beforeCallback-------", req.Host)
	//高妙访问人/ip/uirl/
	return context.WithValue(ctx, "ip", req.Host)
}
