package cache

import (
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type TTL time.Duration

var DefaultTTL = 2 * time.Minute
var NoTTL *TTL = nil

// Cache represents a cache layer.
type Cache[K comparable, V any] struct {
	// mu guards cache.
	mu sync.RWMutex

	// cache is the internal implementation.
	cache *ttlcache.Cache[K, V]
}

// Set sets given key value data to the cache.
func (c *Cache[K, V]) Set(key K, value V, ttl *TTL) {
	c.mu.Lock()
	defer c.mu.Unlock()

	itemTTL := DefaultTTL

	if ttl != nil {
		itemTTL = time.Duration(*ttl)
	}

	if ttl == NoTTL {
		itemTTL = ttlcache.NoTTL
	}

	c.cache.Set(key, value, itemTTL)
}

// Get gets given key value and expiration time from cache.
func (c *Cache[K, V]) Get(key K) (*V, *time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item := c.cache.Get(key)

	if item == nil {
		return nil, nil
	}

	v := item.Value()
	exp := item.ExpiresAt()

	return &v, &exp
}

// New returns a pointer to a Cache struct instance.
func New[K comparable, V any]() *Cache[K, V] {
	// Initializes cache with `defaultTTL`.
	cache := ttlcache.New[K, V](
		ttlcache.WithTTL[K, V](time.Duration(DefaultTTL)),
	)
	go cache.Start()

	c := &Cache[K, V]{
		cache: cache,
	}

	return c
}
