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

func NewCache(interval time.Duration) *cache {
	c := &cache{
		cache: map[string]cacheEntry{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result, ok := c.cache[key]
	if ok {
		return result.val, true
	}
	return []byte{}, false
}

func (c *cache) reapLoop(interval time.Duration) {

}
