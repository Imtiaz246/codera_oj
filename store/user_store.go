package store

import (
	"github.com/imtiaz246/codera_oj/models"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func newUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(u *models.User) error {
	return us.db.Create(u).Error
}

func (us *UserStore) GetUserByUsername(username string, u *models.User) error {
	return us.db.Where("username = ?", username).First(u).Error
}

func (us *UserStore) GetUserByUsernameOrEmail(username, email string, u *models.User) error {
	return us.db.Where("username = ?", username).Or("email = ? AND email IS NOT NULL", email).First(u).Error
}

func (us *UserStore) UpdateUser(u *models.User) error {
	return us.db.Save(u).Error
}
