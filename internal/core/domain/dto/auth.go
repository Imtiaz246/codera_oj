package dto

import (
	"time"
)

type UserRegistration struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginResponse struct {
	User                  *User     `json:"User"`
	AccessToken           string    `json:"AccessToken"`
	AccessTokenExpiresAt  time.Time `json:"AccessTokenExpiresAt"`
	RefreshToken          string    `json:"RefreshToken"`
	RefreshTokenExpiresAt time.Time `json:"RefreshTokenExpiresAt"`
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
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Image        string `json:"image"`
}
