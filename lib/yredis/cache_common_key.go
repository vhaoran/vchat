package yredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/vhaoran/vchat/lib/ylog"
)

func CacheGet(keyPrefix string,
	id interface{},
	expired time.Duration,
//如果 Key值不在缓存中时，如何得到它
	fnNoFound func() (interface{}, error),
//key成功后，如何unmarshal
	fnUnmarshal func(s string) (interface{}, error)) (interface{}, error) {
	//defanlt exp
	exp := time.Second * 3600 * 24
	if expired > 0 {
		exp = expired
	}

	key := keyOfCommon(keyPrefix + fmt.Sprint(id))
	ylog.Debug("redis cache key is:", key)

	doCallbackAndSet := func() (interface{}, error) {
		v, err := fnNoFound()
		if err != nil {
			return nil, err
		}
		//
		var s []byte
		if s, err = json.Marshal(v); err == nil {
			count, err := X.Exists(key).Result()

			_, err = X.Set(key, string(s), exp).Result()
			if err != nil {
				ylog.Error("cache_utils.go->", err)
				return v, nil
			}

			//expired
			if count > 0 {
				X.Expire(key, exp)
			}
		}
		return v, nil
	}

	//
	s, err := X.Get(key).Result()
	// if find
	if err == nil {
		ylog.Debug("redis cache hit,key:", key, " v: ", s)
		return fnUnmarshal(s)
	}

	//if not find
	obj1, err1 := doCallbackAndSet()
	if err1 != nil {
		return nil, err1
	}

	return obj1, nil
}

func keyOfCommon(key string) string {
	return fmt.Sprintf("%s/%s", RedisModuleName, key)
}

func CacheClear(keyPrefix string, keys ...interface{}) {
	l := make([]string, 0)
	for _, v := range keys {
		s := keyOfCommon(keyPrefix + fmt.Sprint(v))
		ylog.Debug("redis clear key:", s)
		l = append(l, s)
	}

	_ = X.Del(l...)
}
