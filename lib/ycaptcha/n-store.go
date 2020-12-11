package ycaptcha

import (
	_ "github.com/dchest/captcha"
	"github.com/vhaoran/vchat/lib/yredis"
	"log"
	"time"
)

type MyStore struct {
	m map[string][]byte
}

func NewMyStore() *MyStore {
	return &MyStore{
		m: make(map[string][]byte),
	}
}

func (r *MyStore) Set(id string, digits []byte) {
	log.Println("------captcha set:", id, ":", digits)
	yredis.X.Set(id, digits, 5*time.Minute)
}

func (r *MyStore) Get(id string, clear bool) (digits []byte) {
	s, err := yredis.X.Get(id).Result()
	if err != nil {
		return nil
	}
	digits = []byte(s)
	log.Println("------captcha  get :", id, ":", digits)
	return
}
