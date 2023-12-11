package pokecache

import (
	// "errors"
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Value     []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

func NewCache(ttl time.Duration) *Cache {
	instance := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}
	go instance.reapLoop(ttl)
	return &instance
}

func (c *Cache) Add(url string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheMap := c.cache
	cacheMap[url] = CacheEntry{
		CreatedAt: time.Now(),
		Value:     data,
	}
	fmt.Println("Add data to cache success with key:", url);
}

func (c *Cache) Get(url string) (CacheEntry, bool) {
	if len(url) == 0{
		return CacheEntry{}, false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheMap := c.cache
	entry, ok := cacheMap[url]
	return entry, ok
}

func (c *Cache) Remove(url string) {
	_, ok := c.cache[url];
	if !ok {
		fmt.Println(url, "Not found key!!")
	} else {
		delete(c.cache, url)
		fmt.Println("delete success key: ", url)
	}
}

func (c *Cache) reapLoop(ttl time.Duration) {
	fmt.Println("Reap loop of cache start", c.cache)
	ticker := time.NewTicker(ttl)

	for {
		select {
		case <-ticker.C:
			c.cleanupExpiredEntries(ttl)
		}
	}
}

func (c *Cache) cleanupExpiredEntries(ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	currentTime := time.Now()
	for key, item := range c.cache {
		timeDifference := currentTime.Sub(item.CreatedAt)
		
		if timeDifference > ttl {
			fmt.Println(key,timeDifference,item.CreatedAt);
			c.Remove(key)
		}
	}
}
