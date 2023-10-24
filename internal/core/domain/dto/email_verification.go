package dto

import "time"

type EmailVerificationInfo struct {
	UserName         string
	VerificationLink string
	ExpirationTime   time.Time
}
