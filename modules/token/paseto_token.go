package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"time"
)

const (
	// ErrInvalidKeySize indicates error for invalidate token key.
	// The chacha20poly1305 algo requires a 32 byte or character key.
	ErrInvalidKeySize = "invalid key size: must be exactly %d characters"

	// Issuer indicates the creator of the token
	Issuer = "coderaOJ.com"
)

// PasetoToken describes a paseto token
type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoToken creates PasetoToken instance
func NewPasetoToken(symmetricKey []byte) (TokenManager, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(ErrInvalidKeySize, chacha20poly1305.KeySize)
	}

	token := &PasetoToken{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}

	return token, nil
}

// CreateToken creates a new token for a specific username and duration
func (pt *PasetoToken) CreateToken(claimsInfo *ClaimsInfo, duration time.Duration) (*TokenInfo, error) {
	pasetoPayload, err := NewPasetoPayload(claimsInfo, duration)
	if err != nil {
		return nil, err
	}

	token, err := pt.paseto.Encrypt(pt.symmetricKey, pasetoPayload, nil)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		Token:   token,
		Payload: pasetoPayload,
	}, nil
}

// VerifyToken verifies if the given token is valid or not.
// And also returns the payload if the token is valid.
func (pt *PasetoToken) VerifyToken(token string) (*paseto.JSONToken, error) {
	pasetoPayload := new(paseto.JSONToken)
	if err := pt.paseto.Decrypt(token, pt.symmetricKey, pasetoPayload, nil); err != nil {
		return nil, fmt.Errorf(ErrInvalidToken, err)
	}

	if time.Now().After(pasetoPayload.Expiration) {
		return nil, fmt.Errorf(ErrExpiredToken)
	}

	return pasetoPayload, nil
}

// NewPasetoPayload creates Payload instance for specific username and duration
func NewPasetoPayload(claimsInfo *ClaimsInfo, duration time.Duration) (*paseto.JSONToken, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	curTime := time.Now()
	pasetoTokenPayload := &paseto.JSONToken{
		Issuer:     Issuer,
		Subject:    "codera oj paseto token",
		Expiration: curTime.Add(duration),
		IssuedAt:   curTime,
		Jti:        tokenID.String(),
	}
	pasetoTokenPayload.Set("username", claimsInfo.Username)
	pasetoTokenPayload.Set("clientIP", claimsInfo.ClientIP)
	pasetoTokenPayload.Set("userAgent", claimsInfo.UserAgent)

	return pasetoTokenPayload, nil
}
