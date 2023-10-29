package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"os"
)

type verifyEmailRepo struct {
	ports.GenericInterface[*auth.VerifyEmail]
	*db.Database
}

var _ ports.VerifyEmailRepoInterface = (*verifyEmailRepo)(nil)

func NewVerifyEmailRepo(d *db.Database) ports.VerifyEmailRepoInterface {
	if err := d.DB.AutoMigrate(auth.VerifyEmail{}); err != nil {
		os.Exit(1)
	}
	return &verifyEmailRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*auth.VerifyEmail](d),
	}
}

func (v *verifyEmailRepo) GetVerifyEmailRecordUsingIdToken(id int64, token string) (*auth.VerifyEmail, error) {
	ve := new(auth.VerifyEmail)
	err := v.Preload("User").Where("id = ? AND token = ?", id, token).First(ve).Error
	if err != nil {
		return nil, err
	}
	return ve, nil
}
