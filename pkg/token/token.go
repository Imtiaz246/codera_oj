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

type TokenManager interface {
	CreateToken(username string, duration time.Duration) (*TokenInfo, error)
	VerifyToken(token string) (*Payload, error)
}

type TokenInfo struct {
	Token   string
	Payload *Payload
}
