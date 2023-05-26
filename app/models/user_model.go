package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	AdminRole = iota
	StaffRole
	UserRole
)

type User struct {
	gorm.Model
	Username         string `gorm:"uniqueIndex;not null"`
	Email            string `gorm:"uniqueIndex,omitempty"`
	KeepEmailPrivate bool   `gorm:"default:1"`
	Password         string `gorm:"not null"`
	Role             uint   `gorm:"default:2"`
	Verified         bool   `gorm:"default:0"`

	FirstName    string
	LastName     string
	Organization string
	Country      string
	City         string
	Image        string
}

func (u *User) HashPassword(plain string) error {
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(h)
	return nil
}

func (u *User) CheckPassword(plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err
}

func (u *User) ExtractEmail() string {
	return u.Email
}
