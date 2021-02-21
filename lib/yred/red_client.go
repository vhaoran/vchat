package yred

import (
	"fmt"
	"github.com/go-redis/redis"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X             *redis.Client
	RedModuleName = ""
)

func InitRedis(cfg yconfig.RedisConfig) error {
	cnt, err := NewRedisClient(cfg)
	if err != nil {
		return err
	}
	X = cnt
	return nil
}

func NewRedisClient(cfg yconfig.RedisConfig) (*redis.Client, error) {
	cnt := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               cfg.Addrs[0],
		Dialer:             nil,
		OnConnect:          nil,
		Password:           cfg.Password,
		DB:                 0,
		MaxRetries:         0,
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
	})

	if err := cnt.Ping().Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cnt, nil
}
