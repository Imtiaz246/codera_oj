package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	models2 "github.com/imtiaz246/codera_oj/app/models"
	"time"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required,min=6"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c *fiber.Ctx, u *models2.User, ve *models2.VerifyEmail, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	if err := u.HashPassword(r.User.Password); err != nil {
		return err
	}
	u.Username = r.User.Username
	if err := ve.GenerateToken(); err != nil {
		return err
	}
	ve.Email = r.User.Email
	ve.ExpirationTime = time.Now().Add(10 * time.Minute)
	return nil
}

type userLoginRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c *fiber.Ctx, u *models2.User, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	if r.User.Username == "" && r.User.Email == "" {
		return fmt.Errorf("username or email must be provided")
	}
	if err := u.HashPassword(r.User.Password); err != nil {
		return err
	}
	u.Email = r.User.Email
	u.Username = r.User.Username

	return nil
}

type userUpdateRequest struct {
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

type updatePasswordRequest struct {
	OldPassword string `json:"old-password" validate:"required,min=6"`
	NewPassword string `json:"new-password" validate:"required,min=6"`
}

type requestedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
