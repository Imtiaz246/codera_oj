package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/models/db"
	"github.com/imtiaz246/codera_oj/modules/cache"
	"github.com/imtiaz246/codera_oj/modules/cronera"
	"log"
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

var SessionCache = cache.NewCache[Session]()

func init() {
	if err := db.MigrateModelTables(Session{}); err != nil {
		panic(err)
	}

	sessionRecords, err := GetAllRecords[*Session]()
	if err != nil {
		panic(err)
	}

	for _, sessionRecord := range sessionRecords {
		if err = SessionCache.Set(sessionRecord.ID.String(), *sessionRecord); err != nil {
			panic(err)
		}
	}
	_, err = cronera.New().Every(1).Day().At("03:00").Do(context.Background(), expiredSessionRemover)
	if err != nil {
		panic(err)
	}
}

func expiredSessionRemover() {
	log.Println("removing expired session started")
	sessionRecords, err := GetAllRecords[*Session]()
	if err != nil {
		log.Println("session record getting error: ", err)
		return
	}

	for _, sessionRecord := range sessionRecords {
		if sessionRecord.ExpiresAt.Before(time.Now()) || sessionRecord.IsBlocked {
			if err = DeleteRecord[*Session](sessionRecord); err != nil {
				log.Println("expired session deletion error: ", err)
			}
			SessionCache.Remove(sessionRecord.ID.String())
		}
	}
}
