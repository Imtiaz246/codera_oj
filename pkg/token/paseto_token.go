package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

const (
	ErrInvalidKeySize = "invalid key size: must be exactly %d characters"
)

// PasetoToken describes a paseto token
type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoToken creates PasetoToken instance
func NewPasetoToken(symmetricKey string) (Token, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(ErrInvalidKeySize, chacha20poly1305.KeySize)
	}

	token := &PasetoToken{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return token, nil
}

// CreateToken creates a new token for a specific username and duration
func (pt *PasetoToken) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return pt.paseto.Encrypt(pt.symmetricKey, payload, nil)
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
