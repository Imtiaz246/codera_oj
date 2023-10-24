package models

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"time"
)

type VerifyEmail struct {
	gorm.Model
	Token          string `gorm:"uniqueIndex;not null"`
	Email          string `gorm:"not null"`
	ExpirationTime time.Time
	IsUsed         bool
	UserId         uint
	User           User
}

func (ve *VerifyEmail) SetVerificationToken() error {
	b := make([]byte, 15)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	ve.Token = hex.EncodeToString(b)
	return nil
}

func (ve *VerifyEmail) IsLinkExpired() error {
	curTime := time.Now()
	if curTime.After(ve.ExpirationTime) {
		return errors.New("link has been expired")
	}
	return nil
}

func (ve *VerifyEmail) IsLinkUsed() bool {
	return ve.IsUsed
}

func (ve *VerifyEmail) ExtractEmail() string {
	return ve.Email
}

func (ve *VerifyEmail) SetExpirationTime() {
	ve.ExpirationTime = time.Now().Add(time.Minute * 15)
}
