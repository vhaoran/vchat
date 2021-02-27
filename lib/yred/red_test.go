package yred

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis"
	"github.com/vhaoran/vchat/lib/ylog"
	"testing"
	"time"
)

func Test_aa(t *testing.T) {
	//
	opt := redis.Options{
		Network:            "",
		Addr:               "192.168.0.99:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         3,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}

	c := redis.NewClient(&opt)
	for ii := 0; ii < 10; ii++ {
		go func() {
			for i := 0; i < 10; i++ {
				k := fmt.Sprint("k_", i)
				v := fmt.Sprint("v_", i)
				{
					_, err := c.Set(k, v, time.Hour).Result()
					if err != nil {
						spew.Dump(err)
					} else {
						fmt.Println("--------set- -ok---")
					}
				}
				{
					s, err := c.Get(k).Result()
					if err != nil {
						spew.Dump(err)
					} else {
						ylog.Debug("-------get- ----", s)
					}
				}
			}
		}()
	}

	ylog.Debug("--------waiting about 20 seconds--- ----")
	time.Sleep(time.Second * 20)
}

func call(c *redis.Client) bool {
	if c == nil {
		return false
	}
	return true
}

func Test_aaa(t *testing.T) {

}
