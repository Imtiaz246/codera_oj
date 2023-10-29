package auth

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	AdminRole Role = "admin"
	StaffRole Role = "staff"
	UserRole  Role = "user"
)

type Role string

type User struct {
	gorm.Model
	Username         string `gorm:"uniqueIndex;not null"`
	Email            string `gorm:"uniqueIndex,omitempty;default:null"`
	KeepEmailPrivate bool   `gorm:"default:1"`
	Password         string `gorm:"not null"`
	Role             Role   `gorm:"default:user"`
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

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) IsAdmin() bool {
	return u.Role == AdminRole
}

func (u *User) IsStaff() bool {
	return u.Role == StaffRole
}

func (u *User) IsUser() bool {
	return u.Role == UserRole
}

func (u *User) ToAPIFormat() *dto.User {
	return &dto.User{
		ID:           int64(u.ID),
		Username:     u.Username,
		Email:        u.Email,
		DisplayName:  u.DisplayName,
		Organization: u.Organization,
		Country:      u.Country,
		City:         u.City,
		Image:        u.Image,
	}
}
