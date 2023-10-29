package ports

import (
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
)

type GenericInterface[T models.ModelFactory] interface {
	GetAllRecords() ([]T, error)
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
	GetVerifyEmailRecordUsingIdToken(id int64, token string) (*models.VerifyEmail, error)
}

type SessionRepoInterface interface {
	GenericInterface[*models.Session]
	GetSessionListOfUser(userID int64) ([]models.Session, error)
	GetSessionByTokenUUID(id uuid.UUID) (*models.Session, error)
}
