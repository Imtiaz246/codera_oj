package ports

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"time"
)

// TokenAdapter is an interface for interacting with token-related logic
type TokenAdapter interface {
	// CreateAccessToken creates a short-lived access token for a user
	CreateAccessToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error)
	// CreateRefreshToken creates a long-lived access token for a user
	CreateRefreshToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error)
	// CreateToken creates a token with claims for a specific duration
	CreateToken(claimsInfo *dto.TokenClaims, duration time.Duration) (*dto.TokenInfo, error)
	// VerifyToken verifies the token and returns the payload
	VerifyToken(token string) (*dto.TokenInfo, error)
}
