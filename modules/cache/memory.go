package cache

import (
	"fmt"
	"sync"
	"time"
)

// Cache is the cache instance to manage session cache
type Cache struct {
	mu         sync.Mutex
	data       map[any]any
	lastAccess time.Time
}

func NewCache() *Cache {
	return &Cache{
		data:       make(map[any]any),
		lastAccess: time.Now(),
	}
}

// lock is used to synchronize the data processing from map
func (c *Cache) lock() {
	c.mu.Lock()
}

// unlock is used to synchronize the data processing from map
func (c *Cache) unlock() {
	c.mu.Unlock()
}

// Get returns the user session value with respect to a key
func (c *Cache) Get(key any) (any, error) {
	c.lock()
	defer c.unlock()

	value, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key: %v, not found", key)
	}

	return value, nil
}

// Set sets a session value with respect to a key
func (c *Cache) Set(key any, value any) error {
	c.lock()
	defer c.unlock()

	for i := 0; i < 3; i++ {
		c.data[key] = value
		if _, ok := c.data[key]; ok {
			return nil
		}
	}
	return fmt.Errorf("internal server problem. Key has not been saved")
}

// Remove deletes a key from SessionCache map
func (c *Cache) Remove(key any) {
	c.lock()
	defer c.unlock()

	delete(c.data, key)
}

// Flush removes all element from the cache
func (c *Cache) Flush() {
	c.lock()
	defer c.unlock()

	for key, _ := range c.data {
		delete(c.data, key)
	}
}
