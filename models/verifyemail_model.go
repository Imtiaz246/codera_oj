package models

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"gorm.io/gorm"
	"strconv"
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

func (ve *VerifyEmail) GenerateToken() error {
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

func (ve *VerifyEmail) GenerateLink() string {
	serverConfig := config.GetServerConfig()
	return serverConfig.Protocol + "://" + serverConfig.Url + "/api/v1/auth/verify-email/" + strconv.FormatUint(uint64(ve.ID), 10) + "/" + ve.Token
}

func (ve *VerifyEmail) ExtractEmail() string {
	return ve.Email
}

func (ve *VerifyEmail) VerifiedUser() *User {
	u := &ve.User
	u.Email = ve.Email
	u.Verified = true

	return u
}

func (ve *VerifyEmail) SetExpirationTime() {
	ve.ExpirationTime = time.Now().Add(time.Minute * 15)
}

func (ve *VerifyEmail) FillEmailVerifierInfo(u *User) error {
	ve.User = *u
	ve.SetExpirationTime()
	if err := ve.GenerateToken(); err != nil {
		return err
	}

	return nil
}
