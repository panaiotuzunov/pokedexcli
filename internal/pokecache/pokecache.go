package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) cache {
	return cache{}
}

func (c cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c cache) Get(key string) ([]byte, bool) {
	result, ok := c.cache[key]
	if ok {
		return result.val, true
	}
	return []byte{}, false
}

func (c cache) reapLoop() {

}
