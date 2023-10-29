package paseto

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"github.com/o1egl/paseto"
	"time"
)

const (
	// ErrInvalidKeySize indicates error for an invalid token key.
	// The chacha20poly1305 algo requires a 32 byte or character key.
	ErrInvalidKeySize = "invalid key size: must be exactly %d characters"

	// Issuer indicates the creator of the token
	Issuer = "coderaOJ.com"
)

// PasetoToken describes a paseto token
type PasetoToken struct {
	paseto *paseto.V2
	config *dto.TokenConfig
}

var _ ports.TokenAdapter = (*PasetoToken)(nil)

// NewPasetoToken creates PasetoToken instance
func NewPasetoToken(tc *dto.TokenConfig) (ports.TokenAdapter, error) {
	if len(tc.Key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(ErrInvalidKeySize, chacha20poly1305.KeySize)
	}

	token := &PasetoToken{
		paseto: paseto.NewV2(),
		config: tc,
	}

	return token, nil
}

// CreateAccessToken creates a short-lived access token for a user
func (pt *PasetoToken) CreateAccessToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error) {
	accessTokenInfo, err := pt.CreateToken(claimsInfo, pt.config.AccessTokenDuration)
	if err != nil {
		return accessTokenInfo, err
	}

	return accessTokenInfo, nil
}

// CreateRefreshToken creates a long-lived access token for a user
func (pt *PasetoToken) CreateRefreshToken(claimsInfo *dto.TokenClaims) (*dto.TokenInfo, error) {
	refreshTokenInfo, err := pt.CreateToken(claimsInfo, pt.config.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}

	return refreshTokenInfo, nil
}

// CreateToken create a new token for a specific duration using claimsInfo
func (pt *PasetoToken) CreateToken(claimsInfo *dto.TokenClaims, duration time.Duration) (*dto.TokenInfo, error) {
	pasetoPayload, err := createPasetoPayload(claimsInfo, duration)
	if err != nil {
		return nil, err
	}

	token, err := pt.paseto.Encrypt(pt.config.Key, pasetoPayload, nil)
	if err != nil {
		return nil, err
	}

	return createTokenInfo(token, pasetoPayload)
}

// VerifyToken verifies if the given token is valid or not.
// And also returns the payload if the token is valid.
func (pt *PasetoToken) VerifyToken(token string) (*dto.TokenInfo, error) {
	pasetoPayload := new(paseto.JSONToken)
	if err := pt.paseto.Decrypt(token, pt.config.Key, pasetoPayload, nil); err != nil {
		return nil, fmt.Errorf("token is not valid: `%v`", err)
	}

	if time.Now().After(pasetoPayload.Expiration) {
		return nil, fmt.Errorf("token is expired")
	}

	return createTokenInfo(token, pasetoPayload)
}
