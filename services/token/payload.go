package token

import (
	"github.com/google/uuid"
	"time"
)

const (
	Issuer = "coderaOJ.com"
)

// Payload contains the payload data for a token
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Iat      time.Time `json:"iat"`
	Exp      time.Time `json:"exp"`
	Iss      string    `json:"iss"`
}

// NewPayload creates Payload instance for specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	curTime := time.Now()
	payload := &Payload{
		ID:       tokenID,
		Username: username,
		Iat:      curTime,
		Exp:      curTime.Add(duration),
		Iss:      Issuer,
	}

	return payload, nil
}

// IsExpired checks the expiration time of a token
func (pl *Payload) IsExpired() bool {
	return time.Now().After(pl.Exp)
}
