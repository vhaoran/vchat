package yredis

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib/yconfig"
)

type Good struct {
	ID   int64
	Name string
}

func (Good) TableName() string {
	return "good"
}

func Test_redis(t *testing.T) {
	url := []string{
		"192.168.0.99:7001",
		"192.168.0.99:7002",
		"192.168.0.99:7003",
		"192.168.0.99:7004",
		"192.168.0.99:7005",
		"192.168.0.99:7006",
	}

	cfg := yconfig.RedisConfig{
		Addrs:              url,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           100,
		MinIdleConns:       10,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
	}

	red, err := NewRedisClient(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	//
	t0 := time.Now()
	h := 100000
	var wg sync.WaitGroup
	wg.Add(h)
	for i := 0; i < h; i++ {
		go func(k int) {
			defer wg.Done()
			if err := red.Set(fmt.Sprint("a_", k), k, time.Hour*100).Err(); err != nil {
				log.Println(err)
			} else {
				if k%1000 == 0 {
					log.Println("OK", k)
				}
			}
		}(i)
	}
	//
	wg.Wait()
	//
	fmt.Println("---------aaa---count--", h, "time:",
		time.Since(t0))
}

func Test_single_cnt(t *testing.T) {
	url := []string{
		"192.168.0.99:6379",
	}

	cfg := yconfig.RedisConfig{
		Addrs:              url,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           100,
		MinIdleConns:       10,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
	}

	red, err := NewRedisClient(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	//
	t0 := time.Now()
	h := 100
	var wg sync.WaitGroup
	wg.Add(h)
	for i := 0; i < h; i++ {
		go func(k int) {
			defer wg.Done()
			if err := red.Set(fmt.Sprint("a_", k), k, time.Hour*1).Err(); err != nil {
				log.Println(err)
			} else {
				if k%1000 == 0 {
					log.Println("OK", k)
				}
			}
		}(i)
	}
	//
	wg.Wait()
	//
	fmt.Println("---------aaa---count--", h, "time:",
		time.Since(t0))

}
