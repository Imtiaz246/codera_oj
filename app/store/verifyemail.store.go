package store

import (
	models2 "github.com/imtiaz246/codera_oj/app/models"
	"gorm.io/gorm"
)

type VerifyEmailStore struct {
	db *gorm.DB
}

func NewVerifyEmailStore(db *gorm.DB) *VerifyEmailStore {
	return &VerifyEmailStore{
		db: db,
	}
}

func (vs *VerifyEmailStore) Create(ve *models2.VerifyEmail, u *models2.User) error {
	ve.User = *u
	ve.GenerateToken()
	return vs.db.Create(ve).Error
}

func (vs *VerifyEmailStore) GetIDToken(id, token string, ve *models2.VerifyEmail) error {
	return vs.db.Preload("User").Where("id = ? AND token = ?", id, token).First(ve).Error
}
