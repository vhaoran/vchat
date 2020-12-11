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
