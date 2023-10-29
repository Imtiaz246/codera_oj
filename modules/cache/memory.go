package cache

import (
	"fmt"
	"sync"
	"time"
)

// memoryCache is the MemoryCache instance to manage session MemoryCache
type memoryCache[T any] struct {
	mu         sync.Mutex
	data       map[string]T
	lastAccess time.Time
}

func NewMemoryCache[T any]() Cache[T] {
	return &memoryCache[T]{
		data:       make(map[string]T),
		lastAccess: time.Now(),
	}
}

// lock is used to synchronize the data processing from map
func (c *memoryCache[T]) lock() {
	c.mu.Lock()
}

// unlock is used to synchronize the data processing from map
func (c *memoryCache[T]) unlock() {
	c.mu.Unlock()
}

// Get returns the user session T with respect to a key
func (c *memoryCache[T]) Get(key string) (*T, error) {
	c.lock()
	defer c.unlock()

	data, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key: %v, not found", key)
	}

	return &data, nil
}

// Set sets a session T with respect to a key
func (c *memoryCache[T]) Set(key string, value T) error {
	c.lock()
	defer c.unlock()

	c.data[key] = value
	if _, ok := c.data[key]; ok {
		return nil
	}

	return fmt.Errorf("internal server problem. Key has not been saved")
}

// Remove deletes a key from SessionCache map
func (c *memoryCache[T]) Remove(key string) {
	c.lock()
	defer c.unlock()

	delete(c.data, key)
}

// Flush removes all element from the MemoryCache
func (c *memoryCache[T]) Flush() {
	c.lock()
	defer c.unlock()

	for key, _ := range c.data {
		delete(c.data, key)
	}
}
