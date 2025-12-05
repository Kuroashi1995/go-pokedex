package pokecache

import(
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	Data	map[string]cacheEntry
	Mu		sync.RWMutex
	ticker	time.Ticker
}
