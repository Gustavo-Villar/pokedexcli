// Package pokecache provides a simple in-memory cache system specifically designed for caching responses from the PokeAPI.
// It leverages mutexes for safe concurrent access and automatically cleans up expired entries.
package pokecache

import (
	"sync"
	"time"
)

// Cache struct defines the structure of the cache, including the cache itself and a mutex for thread-safe operations.
type Cache struct {
	cache map[string]cacheEntry // cache stores the cached data with string keys.
	mux   *sync.Mutex           // mux is used to ensure thread-safe access to the cache.
}

// cacheEntry struct represents a single cache entry, storing the value and the creation time of the entry.
type cacheEntry struct {
	val       []byte    // val is the cached data.
	createdAt time.Time // createdAt is the time when the cache entry was created.
}

// NewCache initializes and returns a new Cache instance with a specified cleanup interval.
// It starts a background goroutine that periodically removes expired cache entries.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval) // Starts the cleanup loop.
	return c
}

// Add inserts a new entry into the cache with the specified key and value.
// It also records the current time as the creation time for the entry.
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock() // Ensures exclusive access to the cache.
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(), // Stores the creation time in UTC.
	}
}

// Get retrieves an entry from the cache by its key.
// It returns the cached value and a boolean indicating whether the key was found in the cache.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock() // Ensures exclusive access to the cache.
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok // Returns the cached value and the presence indicator.
}

// reapLoop is a background goroutine that periodically triggers cache cleanup based on the specified interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // Creates a ticker that emits at the specified interval.
	for range ticker.C {
		c.reap(interval) // Calls the cleanup function at each tick.
	}
}

// reap removes expired cache entries based on their creation time and the specified cleanup interval.
func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock() // Ensures exclusive access to the cache during cleanup.
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval) // Calculates the expiration time.
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) { // Checks if the entry is older than the expiration time.
			delete(c.cache, k) // Removes expired entries.
		}
	}
}
