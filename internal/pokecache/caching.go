package pokecache

import "time"

func NewCache(duration time.Duration) *Cache {
	ticker := time.NewTicker(duration)
	newCache := &Cache{
		Data: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(ticker.C, duration)
	return newCache
}


func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		val:		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	hit, ok := c.Data[key]
	if ok {
		return hit.val, ok
	} else {
		return nil, ok
	}
}

func (c *Cache) reapLoop(ch <-chan time.Time, duration time.Duration) {
	for currentTime := range ch {
		limitTime := currentTime.Add(-duration)
		c.Mu.Lock()
		for k, v := range c.Data {
			if  limitTime.Before(v.createdAt){
				delete(c.Data, k)
			}
		}
		c.Mu.Unlock()
	}
}

func (c *Cache) StopReapLoop() {
	c.ticker.Stop()
}
