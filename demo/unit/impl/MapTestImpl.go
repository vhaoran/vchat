package ctl

import (
	"context"
	"github.com/vhaoran/vchat/demo/unit/intf"
)

type MapTestImpl struct {
}

func (r *MapTestImpl) Exec(ctx context.Context, in *intf.MapTestIn) (*intf.MapTestOut, error) {
	bean := &intf.MapTestOut{
		A:    1,
		B:    "2-string",
		C:    3,
		Code: 200,
		Msg:  "good" + in.S,
		Data: []string{"2222", "333", "44444"},
	}

	//--------ss -----------------------------
	return bean, nil
}
