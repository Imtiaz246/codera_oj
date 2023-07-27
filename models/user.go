package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
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
	Email            string `gorm:"uniqueIndex,omitempty;default:null"`
	KeepEmailPrivate bool   `gorm:"default:1"`
	Password         string `gorm:"not null"`
	Role             uint   `gorm:"default:2"`
	Verified         bool   `gorm:"default:0"`

	DisplayName  string
	Organization string
	Country      string
	City         string
	Image        string
}

func (u *User) HashPassword() error {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(h)
	return nil
}

func (u *User) CheckPassword(plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u *User) ExtractEmail() string {
	return u.Email
}

func init() {
	if err := db.MigrateModelTables(User{}); err != nil {
		panic(err)
	}
}

func GetUserByUsername(username string, u *User) error {
	return db.GetEngine().Where("username = ?", username).First(u).Error
}

func GetUserByUsernameOrEmail(username, email string, u *User) error {
	return db.GetEngine().Where("username = ?", username).Or("email = ? AND email IS NOT NULL", email).First(u).Error
}
