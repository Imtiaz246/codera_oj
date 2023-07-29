package structs

type UserUpdateRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Image        string `json:"image"`
}

type UserUpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
