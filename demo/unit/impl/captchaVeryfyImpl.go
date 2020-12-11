package ctl

import (
	"context"
	"github.com/vhaoran/vchat/demo/unit/intf"
	"github.com/vhaoran/vchat/lib/ycaptcha"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
)

type CaptchaVerifyImpl struct {
}

func (r *CaptchaVerifyImpl) Exec(ctx context.Context, in *intf.CaptchaVerifyIn) (*ykit.Result, error) {
	b := ycaptcha.Verify(in.ID, in.Code)
	log.Println("----------", "param", "------------")
	ylog.DebugDump("in", in)

	return ykit.ROK(b), nil
}
