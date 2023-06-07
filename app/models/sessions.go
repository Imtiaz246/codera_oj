package models

import (
	"github.com/google/uuid"
	"time"
)

type Sessions struct {
	ID        uuid.UUID `gorm:"primarykey"`
	UserId    uint
	User      *User
	UserAgent string
	IsBlocked bool `gorm:"default:0"`
	ClientIP  string
	ExpiresAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;default:null"`
}
