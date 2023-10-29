package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID `gorm:"primarykey"`
	UserID    uint      `gorm:"index"`
	User      *User
	UserAgent string
	IsBlocked bool `gorm:"default:0"`
	ClientIP  string
	ExpiresAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;default:null"`
}

func (s *Session) GetSessionIDString() string {
	return s.ID.String()
}
