package cache

import (
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

var defaultTTL = 10 * time.Minute

// Cache represents a cache layer.
type Cache[K comparable, V any] struct {
	// mu guards cache.
	mu sync.RWMutex

	// cache is the internal implementation.
	cache *ttlcache.Cache[K, V]
}

// Set sets given key value data to the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache.Set(key, value, defaultTTL)
}

// Get gets given key value and expiration time from cache.
func (c *Cache[K, V]) Get(key K) (V, time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item := c.cache.Get(key)

	return item.Value(), item.ExpiresAt()
}

// New returns a pointer to a Cache struct instance.
func New[K comparable, V any]() *Cache[K, V] {
	// Initializes cache with `defaultTTL`.
	cache := ttlcache.New[K, V](
		ttlcache.WithTTL[K, V](defaultTTL),
	)
	go cache.Start()

	c := &Cache[K, V]{
		cache: cache,
	}

	return c
}
