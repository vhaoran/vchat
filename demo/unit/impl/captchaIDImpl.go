package ctl

import (
	"context"
	"github.com/vhaoran/vchat/demo/unit/intf"
	"github.com/vhaoran/vchat/lib/ycaptcha"
	"github.com/vhaoran/vchat/lib/ykit"
)

type CaptchaIDImpl struct {
}

func (r *CaptchaIDImpl) Exec(ctx context.Context, in *intf.CaptchaIDIn) (*ykit.Result, error) {
	return ykit.ROK(ycaptcha.GetCaptchatID()), nil
}
