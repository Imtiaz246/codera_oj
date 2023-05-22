package store

import (
	"github.com/imtiaz246/codera_oj/initializers/db"
	"gorm.io/gorm"
)

type Store struct {
	DB               *gorm.DB
	UserStore        *UserStore
	VerifyEmailStore *VerifyEmailStore
}

func NewStore() (*Store, error) {
	newDB := db.GetDB()
	return &Store{
		DB:               newDB,
		UserStore:        newUserStore(newDB),
		VerifyEmailStore: NewVerifyEmailStore(newDB),
	}, nil
}
