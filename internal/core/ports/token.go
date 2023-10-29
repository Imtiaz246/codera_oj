package ports

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"time"
)

// TokenAdapter is an interface for managing tokens.
type TokenAdapter interface {
	// CreateAccessToken generates a short-lived user access token based on provided token claims.
	CreateAccessToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error)

	// CreateRefreshToken creates a long-lived user refresh token based on provided token claims.
	CreateRefreshToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error)

	// CreateToken generates a token with specific details and duration based on token claims.
	CreateToken(claimsInfo *dto.TokenClaims, duration time.Duration) (*dto.TokenInfo, error)

	// VerifyToken validates and decodes a token, returns its associated information.
	VerifyToken(token string) (*dto.TokenInfo, error)
}
