package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserRegistration struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Username  string `json:"username"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"required,min=6"`
	ClientIP  string `json:"clientIP" validate:"required"`
	UserAgent string `json:"userAgent" validate:"required"`
}

type UserLoginResponse struct {
	User             *User     `json:"User"`
	AccessTokenInfo  TokenInfo `json:"AccessTokenInfo"`
	RefreshTokenInfo TokenInfo `json:"RefreshTokenInfo"`
}

type RequestedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RenewTokenResponse struct {
	User                 *User     `json:"user"`
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Image        string `json:"image"`
}

type EmailVerificationInfo struct {
	UserName         string
	VerificationLink string
	ExpirationTime   time.Time
}

type TokenClaims struct {
	Username  string
	ClientIP  string
	UserAgent string
}

type TokenInfo struct {
	Token      string
	ID         uuid.UUID
	Expiration time.Time
	IssuedAt   time.Time
	TokenClaims
}
