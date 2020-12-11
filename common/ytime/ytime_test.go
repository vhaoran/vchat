package ytime

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func Test_cur_test(t *testing.T) {
	now := time.Now()
	//
	//
	spew.Dump("prior:  ", PriorMonth(now))
	spew.Dump("cur:  ", CurrentMonth(now))
	spew.Dump("next:  ", NextMonth(now))
}
