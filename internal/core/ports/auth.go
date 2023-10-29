package ports

import (
	"context"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
)

// AuthService represents an interface for managing user authentication and authorization.
type AuthService interface {
	// SignUp creates user records and sends an email verification mail for email confirmation.
	SignUp(ctx context.Context, registrationData *dto.UserRegistration) error

	// Login authenticates a user by their email or username and password, returning access and refresh tokens.
	Login(ctx context.Context, loginData dto.UserLogin) (*dto.UserLoginResponse, error)

	// VerifyEmail verifies a user's email using a unique ID and verification token obtained from an email link.
	VerifyEmail(ctx context.Context, id int64, token string) error

	// RenewToken renews an authentication token using a refresh token, and it provides a new, short-lived access token.
	RenewToken(ctx context.Context, token, reqIP, reqUserAgent string) (*dto.TokenInfo, error)
}
