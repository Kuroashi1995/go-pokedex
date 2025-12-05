package pokecache_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

func TestAddCache(t *testing.T) {
	cache := pokecache.NewCache(time.Second * 5)
	defer cache.StopReapLoop()
	// test normal add
	cache.Add("https://example1.com", []byte("testdata"))
	if _, ok := cache.Get("https://example1.com"); !ok {
		t.Errorf("Cache should exist")
		return
	}
	// test rewrite
	cache.Add("https://example1.com", []byte("rewrite"))
	val, ok := cache.Get("https://example1.com")
	if !ok {
		t.Errorf("Cache should exist")
		return
	}
	test := []byte("rewrite")
	if !bytes.Equal(val, test) {
		t.Errorf("Cache rewrite failed, values are different")
		return
	}
}
func TestReapLoop(t *testing.T) {
	// define table struct
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond

	cache := pokecache.NewCache(baseTime)
	defer cache.StopReapLoop()
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Expected to find the key")
		return
	}

	time.Sleep(waitTime)
	
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Expected not to find the key")
		return
	}
	
}
