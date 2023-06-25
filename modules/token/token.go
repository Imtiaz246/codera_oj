package token

import (
	"github.com/o1egl/paseto"
	"time"
)

const (
	// ErrInvalidToken indicates the error of token invalidation
	ErrInvalidToken = "token is invalid: %v"

	// ErrExpiredToken indicated the error of token expiration
	ErrExpiredToken = "token is expired"
)

// TokenManager is the interface for managing tokens
type TokenManager interface {
	CreateToken(claimsInfo *ClaimsInfo, duration time.Duration) (*TokenInfo, error)
	VerifyToken(token string) (*paseto.JSONToken, error)
}

// TokenInfo holds the information related to a token
type TokenInfo struct {
	Token   string
	Payload *paseto.JSONToken
}

// ClaimsInfo holds the claims information for a user
type ClaimsInfo struct {
	Username  string
	ClientIP  string
	UserAgent string
}
