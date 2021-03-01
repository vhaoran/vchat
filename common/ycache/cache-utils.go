package ycache

import (
	"github.com/goburrow/cache"
	"time"
)

func NewCache(loadFunc cache.LoaderFunc) cache.LoadingCache {
	// Create a new cache
	c := cache.NewLoadingCache(loadFunc,
		cache.WithMaximumSize(10000),
		cache.WithExpireAfterAccess(60*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)
	return c
}

func NewCacheExpire(length int, expire time.Duration) cache.Cache {
	// Create a new cache
	if length < 100 {
		length = 100
	}
	if expire < time.Second {
		expire = time.Second
	}

	return cache.New(
		cache.WithMaximumSize(length),
		cache.WithExpireAfterAccess(expire),
		cache.WithRefreshAfterWrite(expire),
	)
}
