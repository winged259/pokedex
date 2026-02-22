package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	mu sync.RWMutex
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache() *cache {
	return &cache{
		entry: make(map[string]cacheEntry),
	}
}

func (c *cache) Add(key string, val []byte) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entry[key]
	if exists {
		return entry.val, true
	}
	
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	return val, false
}