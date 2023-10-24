package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
)

type userRepo struct {
	GenericInterface[*models.User]
	*db.Database
}

var _ UserRepoInterface = (*userRepo)(nil)

func NewUserRepo(d *db.Database) UserRepoInterface {
	return userRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*models.User](d),
	}
}

func (ur userRepo) GetUserByUsernameOrEmail(username, email string) (*models.User, error) {
	u := new(models.User)
	if err := ur.DB.Where("username = ?", username).Or("email = ? AND email IS NOT NULL", email).First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
