package yetcd

import (
	"context"
	"testing"

	"github.com/coreos/etcd/clientv3"
)

func Test_etcd_new(t *testing.T) {
	// initial etcd v3 client
	cnt, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		panic(err)
	}
	cnt.Put(context.Background(), "key", "value")
}
