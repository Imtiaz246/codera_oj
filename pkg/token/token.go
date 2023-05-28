package token

import (
	"time"
)

const (
	// ErrInvalidToken indicates the error of token invalidation
	ErrInvalidToken = "token is invalid: %v"

	// ErrExpiredToken indicated the error of token expiration
	ErrExpiredToken = "token is expired"
)

type Token interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
