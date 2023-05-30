package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

const (
	// ErrInvalidKeySize indicates error for invalidate token key.
	// The chacha20poly1305 algo requires a 32 byte or character key.
	ErrInvalidKeySize = "invalid key size: must be exactly %d characters"
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
func (pt *PasetoToken) CreateToken(username string, duration time.Duration) (*TokenInfo, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return nil, err
	}

	token, err := pt.paseto.Encrypt(pt.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		Token:   token,
		Payload: payload,
	}, nil
}

// VerifyToken verifies if the given token is valid or not
func (pt *PasetoToken) VerifyToken(token string) (*Payload, error) {
	payload := new(Payload)
	if err := pt.paseto.Decrypt(token, pt.symmetricKey, payload, nil); err != nil {
		return nil, fmt.Errorf(ErrInvalidToken, err)
	}

	if ok := payload.IsExpired(); !ok {
		return nil, fmt.Errorf(ErrExpiredToken)
	}

	return nil, nil
}
