package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

var (
	mu    sync.Mutex
	Cache = make(map[string]CacheItem)
)

func Get(key string) interface{} {
	mu.Lock()
	defer mu.Unlock()

	item, found := Cache[key]
	if !found {
		return nil
	}

	// Check if the item has expired, unless it's a permanent item
	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		delete(Cache, key)
		return nil
	}

	return item.Value
}

func Set(key string, value interface{}, duration time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	var expiration int64
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	Cache[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}
