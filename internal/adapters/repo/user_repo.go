package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"os"
)

type userRepo struct {
	ports.GenericInterface[*models.User]
	*db.Database
}

var _ ports.UserRepoInterface = (*userRepo)(nil)

func NewUserRepo(d *db.Database) ports.UserRepoInterface {
	if err := d.DB.AutoMigrate(models.User{}); err != nil {
		os.Exit(1)
	}
	return &userRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*models.User](d),
	}
}

func (ur *userRepo) GetUserByUsernameOrEmail(username, email string) (*models.User, error) {
	u := new(models.User)
	if err := ur.DB.Where("username = ?", username).Or("email = ? AND email IS NOT NULL", email).First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
