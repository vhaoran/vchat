package ymid

import (
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/vhaoran/vchat/lib/ylog"
)

//内置中间件
func MidCommon(ep endpoint.Endpoint) endpoint.Endpoint {
	t0 := time.Now()
	defer func() {
		ylog.Debug("#####exec time(ms):", time.Since(t0).Milliseconds())
	}()

	return ep
}
