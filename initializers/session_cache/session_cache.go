package session_cache

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/db"
)

// sessionCache is the global cache for managing user sessions
var sessionCache map[string]*models.Sessions

// LoadSessionCache loads the user session from the persistent database
func LoadSessionCache() error {
	if ok := db.IsDBInitialized(); !ok {
		return fmt.Errorf("db is not initialized")
	}
	sessionCache = make(map[string]*models.Sessions)

	database := db.GetDB()
	var sessionRecords []models.Sessions
	if err := database.Find(&sessionRecords).Error; err != nil {
		return err
	}

	for _, record := range sessionRecords {
		recordCopy := record
		sessionCache[record.ID.String()] = &recordCopy
	}

	return nil
}

// Get returns the user session value with respect to a key
func Get(key string) (*models.Sessions, error) {
	value, ok := sessionCache[key]
	if !ok {
		return nil, fmt.Errorf("key: %v, not found", key)
	}

	return value, nil
}

// Set sets a session value with respect to a key
func Set(key string, value *models.Sessions) error {
	for i := 0; i < 3; i++ {
		sessionCache[key] = value
		if _, ok := sessionCache[key]; ok {
			return nil
		}
	}
	return fmt.Errorf("internal server problem. Key has not been saved")
}
