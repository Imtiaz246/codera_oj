package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"os"
)

type userRepo struct {
	ports.GenericInterface[*auth.User]
	*db.Database
}

var _ ports.UserRepoInterface = (*userRepo)(nil)

func NewUserRepo(d *db.Database) ports.UserRepoInterface {
	if err := d.DB.AutoMigrate(auth.User{}); err != nil {
		os.Exit(1)
	}
	return &userRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*auth.User](d),
	}
}

func (ur *userRepo) GetUserByUsernameOrEmail(username, email string) (*auth.User, error) {
	u := new(auth.User)
	if err := ur.DB.Where("username = ?", username).Or("email = ? AND email IS NOT NULL", email).First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
