package cache

import (
	"fmt"
	"sync"
	"time"
)

// Cache is the cache instance to manage session cache
type Cache[VT any] struct {
	mu         sync.Mutex
	data       map[string]VT
	lastAccess time.Time
}

func NewCache[VT any]() *Cache[VT] {
	return &Cache[VT]{
		data:       make(map[string]VT),
		lastAccess: time.Now(),
	}
}

// lock is used to synchronize the data processing from map
func (c *Cache[VT]) lock() {
	c.mu.Lock()
}

// unlock is used to synchronize the data processing from map
func (c *Cache[VT]) unlock() {
	c.mu.Unlock()
}

// Get returns the user session VT with respect to a key
func (c *Cache[VT]) Get(key string) (*VT, error) {
	c.lock()
	defer c.unlock()

	data, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key: %v, not found", key)
	}

	return &data, nil
}

// Set sets a session VT with respect to a key
func (c *Cache[VT]) Set(key string, value VT) error {
	c.lock()
	defer c.unlock()

	c.data[key] = value
	if _, ok := c.data[key]; ok {
		return nil
	}

	return fmt.Errorf("internal server problem. Key has not been saved")
}

// Remove deletes a key from SessionCache map
func (c *Cache[VT]) Remove(key string) {
	c.lock()
	defer c.unlock()

	delete(c.data, key)
}

// Flush removes all element from the cache
func (c *Cache[VT]) Flush() {
	c.lock()
	defer c.unlock()

	for key, _ := range c.data {
		delete(c.data, key)
	}
}
