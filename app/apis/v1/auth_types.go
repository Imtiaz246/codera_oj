package v1

import "github.com/imtiaz246/codera_oj/app/models"

var (
	UserSuccessfulRegistrationMessage = struct {
		Message string `json:"message"`
	}{
		Message: "Account registered successfully. Please verify your email to add the email to your profile.",
	}

	EmailSuccessfulVerificationMessage = struct {
		Message string `json:"message"`
	}{
		Message: "Email verified successfully.",
	}
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required,min=6"`
	} `json:"user"`
}

type UserLoginRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

type UserUpdateRequest struct {
	User struct {
		Email        string `json:"email"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Organization string `json:"organization"`
		Country      string `json:"country"`
		City         string `json:"city"`
		Image        string `json:"image"`
	} `json:"user"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old-password" validate:"required,min=6"`
	NewPassword string `json:"new-password" validate:"required,min=6"`
}

type RequestedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	User struct {
		ID           uint   `json:"id"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Organization string `json:"organization"`
		Country      string `json:"country"`
		City         string `json:"city"`
		Image        string `json:"image"`
	} `json:"user"`
}

func NewUserResponse(u *models.User) *UserResponse {
	r := new(UserResponse)
	r.User.ID = u.ID
	r.User.City = u.City
	r.User.Image = u.Image
	r.User.Country = u.Country
	r.User.Username = u.Username
	r.User.LastName = u.LastName
	r.User.FirstName = u.FirstName
	r.User.Organization = u.Organization

	if u.KeepEmailPrivate == false {
		r.User.Email = u.Email
	}

	return r
}
