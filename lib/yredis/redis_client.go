package yredis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X               *redis.ClusterClient
	RedisModuleName = ""
)

func InitRedis(cfg yconfig.RedisConfig) error {
	cnt, err := NewRedisClient(cfg)
	if err != nil {
		return err
	}
	X = cnt
	return nil
}

func NewRedisClient(cfg yconfig.RedisConfig) (*redis.ClusterClient, error) {
	cnt := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:          cfg.Addrs,
		MaxRedirects:   cfg.MaxRedirects,
		ReadOnly:       false,
		RouteByLatency: false,
		RouteRandomly:  false,
		ClusterSlots:   nil,
		OnNewNode:      nil,
		//Dialer:             nil,
		OnConnect:          nil,
		Password:           cfg.Password,
		MaxRetries:         cfg.MaxRedirects,
		MinRetryBackoff:    cfg.MinRetryBackoff,
		MaxRetryBackoff:    cfg.MaxRetryBackoff,
		DialTimeout:        cfg.DialTimeout * time.Second,
		ReadTimeout:        cfg.ReadTimeout * time.Second,
		WriteTimeout:       cfg.WriteTimeout * time.Second,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       cfg.MinIdleConns,
		MaxConnAge:         cfg.MaxConnAge * time.Second,
		PoolTimeout:        cfg.PoolTimeout * time.Second,
		IdleTimeout:        cfg.IdleTimeout * time.Second,
		IdleCheckFrequency: cfg.IdleCheckFrequency * time.Second,
		TLSConfig:          nil,
	})

	if err := cnt.Ping().Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cnt, nil
}
