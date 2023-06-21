package session

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/modules/cache"
	"gorm.io/gorm"
	"log"
)

// Cache is the cache instance to manage session cache
type Cache struct {
	*cache.Cache
}

// initSessionCache loads session data from database and returns SessionCache instance
func initSessionCache(db *gorm.DB) (*Cache, error) {
	sc := &Cache{
		cache.NewCache(),
	}

	if err := sc.loadSessionCache(db); err != nil {
		return nil, err
	}

	return sc, nil
}

// loadSessionCache loads the user session from the persistent database
func (sc *Cache) loadSessionCache(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("db is not initialized")
	}

	var sessionRecords []models.Session
	if err := db.Find(&sessionRecords).Error; err != nil {
		return err
	}

	for _, record := range sessionRecords {
		recordCopy := record
		if err := sc.Set(record.ID.String(), &recordCopy); err != nil {
			log.Printf("failed to set key `%v` on session cache", record.ID)
		}
	}

	return nil
}
