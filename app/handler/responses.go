package handler

import (
	"github.com/imtiaz246/codera_oj/app/models"
)

type userResponse struct {
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

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
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
