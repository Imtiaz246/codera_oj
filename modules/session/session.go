package session

import "gorm.io/gorm"

func New(db *gorm.DB) (*Cache, error) {
	sessionCache, err := initSessionCache(db)
	if err != nil {
		return nil, err
	}

	return sessionCache, nil
}
