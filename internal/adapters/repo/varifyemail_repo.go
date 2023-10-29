package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"os"
)

type verifyEmailRepo struct {
	ports.GenericInterface[*models.VerifyEmail]
	*db.Database
}

var _ ports.VerifyEmailRepoInterface = (*verifyEmailRepo)(nil)

func NewVerifyEmailRepo(d *db.Database) ports.VerifyEmailRepoInterface {
	if err := d.DB.AutoMigrate(models.VerifyEmail{}); err != nil {
		os.Exit(1)
	}
	return &verifyEmailRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*models.VerifyEmail](d),
	}
}

func (v *verifyEmailRepo) GetVerifyEmailRecordUsingIdToken(id int64, token string) (*models.VerifyEmail, error) {
	ve := new(models.VerifyEmail)
	err := v.Preload("User").Where("id = ? AND token = ?", id, token).First(ve).Error
	if err != nil {
		return nil, err
	}
	return ve, nil
}
