package ports

import (
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
)

// GenericInterface is an interface for performing generic database operations on a specific model type.
type GenericInterface[T models.ModelFactory] interface {
	// GetAllRecords retrieves all records of the specified model type.
	GetAllRecords() ([]T, error)

	// GetRecordByModel retrieves a record by the model type with optional preloads.
	GetRecordByModel(t T, preloads ...string) error

	// GetRecordByID retrieves a record by its unique identifier.
	GetRecordByID(id uint) (T, error)

	// CreateRecord creates a new record.
	CreateRecord(t T) error

	// UpdateRecord updates an existing record.
	UpdateRecord(t T) error

	// DeleteRecord deletes a record.
	DeleteRecord(t T) error

	// DeleteRecordByID deletes a record by its unique identifier.
	DeleteRecordByID(id uint) (T, error)

	// GetRecordByExpression retrieves a record based on an expression.
	GetRecordByExpression(query any, args ...any) (T, error)
}

// UserRepoInterface is an interface for performing user-related database operations.
type UserRepoInterface interface {
	GenericInterface[*auth.User]

	// GetUserByUsernameOrEmail retrieves a user by username or email.
	GetUserByUsernameOrEmail(username, email string) (*auth.User, error)
}

// VerifyEmailRepoInterface is an interface for performing email verification-related database operations.
type VerifyEmailRepoInterface interface {
	GenericInterface[*auth.VerifyEmail]

	// GetVerifyEmailRecordUsingIdToken retrieves a verification record using an identifier and token.
	GetVerifyEmailRecordUsingIdToken(id int64, token string) (*auth.VerifyEmail, error)
}

// SessionRepoInterface is an interface for performing session-related database operations.
type SessionRepoInterface interface {
	GenericInterface[*auth.Session]

	// GetSessionListOfUser retrieves a list of sessions for a user.
	GetSessionListOfUser(userID int64) ([]auth.Session, error)

	// GetSessionByTokenUUID retrieves a session by its token UUID.
	GetSessionByTokenUUID(id uuid.UUID) (*auth.Session, error)
}

// ProblemRepoInterface is an interface for performing problem-related database operations.
type ProblemRepoInterface interface {
	GenericInterface[*problem.Problem]
}

// TagRepoInterface is an interface for performing tag-related database operations.
type TagRepoInterface interface {
	GenericInterface[*problem.Tag]
}

// ProblemTagRepoInterface is an interface for performing problem tag-related database operations.
type ProblemTagRepoInterface interface {
	GenericInterface[*problem.ProblemTag]
}

// DiscussionRepoInterface is an interface for performing discussion-related database operations.
type DiscussionRepoInterface interface {
	GenericInterface[*problem.Discussion]
}

// ChangeLogRepoInterface is an interface for performing change log-related database operations.
type ChangeLogRepoInterface interface {
	GenericInterface[*problem.ChangeLog]
}

// LanguageRepoInterface is an interface for performing language-related database operations.
type LanguageRepoInterface interface {
	GenericInterface[*problem.Language]
}

// ShareRepoInterface is an interface for performing share-related database operations.
type ShareRepoInterface interface {
	GenericInterface[*problem.Share]
}

// SolutionRepoInterface is an interface for performing solution-related database operations.
type SolutionRepoInterface interface {
	GenericInterface[*problem.Solution]
}

// DatasetRepoInterface is an interface for performing dataset-related database operations.
type DatasetRepoInterface interface {
	GenericInterface[*problem.Dataset]
}
