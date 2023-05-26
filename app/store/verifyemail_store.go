package store

import (
	"github.com/imtiaz246/codera_oj/app/models"
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

func (vs *VerifyEmailStore) Create(ve *models.VerifyEmail) error {
	return vs.db.Create(ve).Error
}

func (vs *VerifyEmailStore) GetIDToken(id, token string, ve *models.VerifyEmail) error {
	return vs.db.Preload("User").Where("id = ? AND token = ?", id, token).First(ve).Error
}

func (vs *VerifyEmailStore) Update(ve *models.VerifyEmail) error {
	return vs.db.Save(ve).Error
}
