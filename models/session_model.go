package models

import (
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/models/db"
	"github.com/imtiaz246/codera_oj/modules/cache"
	"time"
)

type Session struct {
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

var SessionCache *cache.Cache[Session]

func init() {
	if err := db.MigrateModelTables(Session{}); err != nil {
		panic(err)
	}

	// Initialize session cache
	SessionCache = cache.NewCache[Session]()

	sessionRecords, err := GetAllRecords[*Session]()
	if err != nil {
		panic(err)
	}

	for _, sessionRecord := range sessionRecords {
		if err = SessionCache.Set(sessionRecord.ID.String(), *sessionRecord); err != nil {
			panic(err)
		}
	}

	// todo: add expired session remover cron job
}
