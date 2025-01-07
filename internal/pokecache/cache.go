package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	locations map[string]cacheEntry
	mt        sync.RWMutex
	interval  time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(configs CacheConfig) *Cache {
	if configs.Interval == 0 {
		log.Printf("Could not load cache cleaning interval. Using 1 minute")
	}

	newCache := Cache{
		locations: make(map[string]cacheEntry),
		mt:        sync.RWMutex{},
		interval:  time.Duration(configs.Interval) * time.Millisecond,
	}

	go (&newCache).reapLoop()

	return &newCache
}

func (c *Cache) Add(location string, entry []byte) {
	c.mt.Lock()
	defer c.mt.Unlock()

	c.locations[location] = cacheEntry{
		createdAt: time.Now(),
		val:       entry,
	}
}

func (c *Cache) Get(location string) ([]byte, bool) {
	c.mt.RLock()
	defer c.mt.RUnlock()

	loc, ok := c.locations[location]
	return loc.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		t := <-ticker.C
		c.clearOldEntries(t)
	}
}

func (c *Cache) clearOldEntries(now time.Time) {
	c.mt.Lock()
	defer c.mt.Unlock()

	for k, v := range c.locations {
		if now.Sub(v.createdAt) > c.interval {
			delete(c.locations, k)
		}
	}
}
