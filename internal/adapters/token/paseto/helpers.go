package paseto

import (
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/o1egl/paseto"
	"time"
)

// createPasetoPayload creates Payload instance for specific username and duration
func createPasetoPayload(claimsInfo *dto.TokenClaims, duration time.Duration) (*paseto.JSONToken, error) {
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

// createTokenInfo create tokenInfo instance
func createTokenInfo(token string, pasetoPayload *paseto.JSONToken) (*dto.TokenInfo, error) {
	tokenID, err := uuid.Parse(pasetoPayload.Jti)
	if err != nil {
		return nil, err
	}
	return &dto.TokenInfo{
		TokenClaims: dto.TokenClaims{
			Username:  pasetoPayload.Get("username"),
			ClientIP:  pasetoPayload.Get("clientIP"),
			UserAgent: pasetoPayload.Get("userAgent"),
		},
		Token:      token,
		ID:         tokenID,
		Expiration: pasetoPayload.Expiration,
		IssuedAt:   pasetoPayload.IssuedAt,
	}, nil
}
