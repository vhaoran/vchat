package ctl

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/vhaoran/vchat/demo/unit/intf"
	"github.com/vhaoran/vchat/lib/ykit"
	"log"
)
/*
curl -X POST  \
-H 'Jwt:test/1' \
-H 'Content-Type:application/json' \
-d '{"page_no":1,"rows_per_page":2,"where":[ {"uid":1} ]   }' \
localhost:9999/api/pb


*/



type PBImpl struct {
}

func (r *PBImpl) Exec(ctx context.Context, in *intf.PBIn) (*ykit.Result, error) {
	log.Println("------PBImpl) Fn----", "------------")
	spew.Dump(in)
	log.Println("------end----", "------------")
	return ykit.ROK(in), nil
}
