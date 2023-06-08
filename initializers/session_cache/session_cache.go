package session_cache

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/initializers/db"
	"github.com/imtiaz246/codera_oj/models"
	"log"
	"sync"
)

// sessionCache is the global cache for managing user sessions
type sessionCache struct {
	mu   sync.Mutex
	data map[string]*models.Sessions
}

// lock is used to synchronize the data processing from map
func (sc *sessionCache) lock() {
	sc.mu.Lock()
}

// unlock is used to synchronize the data processing from map
func (sc *sessionCache) unlock() {
	sc.mu.Unlock()
}

// cache is the sessionCache instance
var cache *sessionCache

// LoadSessionCache loads the user session from the persistent database
func LoadSessionCache() error {
	if ok := db.IsDBInitialized(); !ok {
		return fmt.Errorf("db is not initialized")
	}
	cache = &sessionCache{
		data: make(map[string]*models.Sessions),
	}

	database := db.GetDB()
	var sessionRecords []models.Sessions
	if err := database.Find(&sessionRecords).Error; err != nil {
		return err
	}

	for _, record := range sessionRecords {
		recordCopy := record
		if err := Set(record.ID.String(), &recordCopy); err != nil {
			log.Printf("failed to set key `%v` on session cache", record.ID)
		}
	}

	return nil
}

// Get returns the user session value with respect to a key
func Get(key string) (*models.Sessions, error) {
	cache.lock()
	defer cache.unlock()

	value, ok := cache.data[key]
	if !ok {
		return nil, fmt.Errorf("key: %v, not found", key)
	}

	return value, nil
}

// Set sets a session value with respect to a key
func Set(key string, value *models.Sessions) error {
	cache.lock()
	defer cache.unlock()

	for i := 0; i < 3; i++ {
		cache.data[key] = value
		if _, ok := cache.data[key]; ok {
			return nil
		}
	}
	return fmt.Errorf("internal server problem. Key has not been saved")
}

// DeleteFromSessionCache deletes a key from sessionCache map
func DeleteFromSessionCache(key string) {
	cache.lock()
	defer cache.unlock()
	delete(cache.data, key)
}
