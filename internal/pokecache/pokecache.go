package pokecache

import "time"

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(t time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
	}
	go c.purgeLoop(t)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheExists, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return cacheExists.val, ok
}

func (c *Cache) purge(t time.Duration) {
	fiveMinOld := time.Now().UTC().Add(-t)
	for k, v := range c.cache {
		if v.createdAt.Before(fiveMinOld) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) purgeLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.purge(t)
	}
}
