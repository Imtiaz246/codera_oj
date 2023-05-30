package v1

import (
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/services/token"
	"time"
)

var (
	UserSuccessfulRegistrationResponse = struct {
		Message string `json:"message"`
	}{
		Message: "Account registered successfully. Please verify your email to add the email to your profile.",
	}

	EmailSuccessfulVerificationResponse = struct {
		Message string `json:"message"`
	}{
		Message: "Email verified successfully.",
	}
)

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginResponse struct {
	User                  *UserResponse `json:"User"`
	AccessToken           string        `json:"AccessToken"`
	AccessTokenExpiresAt  time.Time     `json:"AccessTokenExpiresAt"`
	RefreshToken          string        `json:"RefreshToken"`
	RefreshTokenExpiresAt time.Time     `json:"RefreshTokenExpiresAt"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type RequestedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Image        string `json:"image"`
}

func NewUserResponse(u *models.User) *UserResponse {
	r := &UserResponse{
		ID:           u.ID,
		City:         u.City,
		Image:        u.Image,
		Country:      u.Country,
		Username:     u.Username,
		LastName:     u.LastName,
		FirstName:    u.FirstName,
		Organization: u.Organization,
	}
	if u.KeepEmailPrivate == false {
		r.Email = u.Email
	}

	return r
}

func NewLoginResponse(u *models.User, atInfo, rtInfo *token.TokenInfo) *UserLoginResponse {
	return &UserLoginResponse{
		User:                  NewUserResponse(u),
		AccessToken:           atInfo.Token,
		AccessTokenExpiresAt:  atInfo.Payload.Exp,
		RefreshToken:          rtInfo.Token,
		RefreshTokenExpiresAt: rtInfo.Payload.Exp,
	}
}
