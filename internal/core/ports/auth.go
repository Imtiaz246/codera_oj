package ports

import (
	"context"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
)

// AuthService is an interface to interact with the user authentication related business logic
type AuthService interface {
	// SignUp create records for a user and send email verification mail for verifying email
	SignUp(ctx context.Context, registrationData *dto.UserRegistration) error
	// Login authenticates a user email or user and password and returns access and refresh token
	Login(ctx context.Context, loginData dto.UserLogin) (*dto.UserLoginResponse, error)
	// VerifyEmail verifies an email for a user form a link with id and token
	VerifyEmail(ctx context.Context, id int64, token string) error
	// RenewToken renews token using refresh token and returns new short-lived access token
	RenewToken(ctx context.Context, token, reqIP, reqUserAgent string) (*dto.TokenInfo, error)
}
