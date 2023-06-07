package store

import (
	"github.com/imtiaz246/codera_oj/app/models"
	"gorm.io/gorm"
)

type SessionStore struct {
	db *gorm.DB
}

func newSessionStore(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (ss *SessionStore) Create(s *models.Sessions) error {
	return ss.db.Create(s).Error
}

func (ss *SessionStore) GetBySessionID(id string, s *models.Sessions) error {
	return ss.db.Where("ID = ?", id).First(s).Error
}
