package repo

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
)

type GenericInterface[T models.ModelFactory] interface {
	GetAllRecords(t T) ([]T, error)
	GetRecordByModel(t T, preloads ...string) error
	GetRecordByID(id int64) (T, error)
	CreateRecord(t T) error
	UpdateRecord(t T) error
	DeleteRecord(t T) error
	DeleteRecordByID(id int64) (T, error)
	GetRecordByExpression(query any, args ...any) (T, error)
}

type UserRepoInterface interface {
	GenericInterface[*models.User]
	GetUserByUsernameOrEmail(username, email string) (*models.User, error)
}

type VerifyEmailRepoInterface interface {
	GenericInterface[*models.VerifyEmail]
}
