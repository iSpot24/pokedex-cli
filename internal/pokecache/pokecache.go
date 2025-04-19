package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		created_at: time.Now(),
		val:        val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]

	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	for k, v := range c.data {
		if time.Since(v.created_at) >= interval {
			delete(c.data, k)
		}
	}
	c.mu.Unlock()
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

type cacheEntry struct {
	created_at time.Time
	val        []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}
