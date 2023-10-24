package port

import (
	"context"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"time"
)

// TokenService is an interface for interacting with token-related business logic
type TokenService interface {
	// CreateToken creates a new token for a given user
	CreateToken(user *models.User, duration time.Duration) (string, error)
	// VerifyToken verifies the token and returns the payload
	//VerifyToken(token string) (*domain.TokenPayload, error)
}

// AuthService is an interface to interact with the user authentication related business logic
type AuthService interface {
	// SignUp create records for a user and send email verification mail for verifying email
	SignUp(ctx context.Context, registrationData *dto.UserRegistration) error
	// Login authenticates a user email or user and password and returns access and refresh token
	Login(ctx context.Context, loginData dto.UserLogin) (*dto.UserLoginResponse, error)
	// VerifyEmail verifies an email for a user form a link with id and token
	VerifyEmail(ctx context.Context, id int64, token string) error
	// RenewToken renews token using refresh token and returns new short-lived access token
	RenewToken(ctx context.Context, refreshToken string) (string, error)
}
