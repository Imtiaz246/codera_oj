package handler

import (
	"github.com/gofiber/fiber/v2"
	apiv1 "github.com/imtiaz246/codera_oj/app/api/v1"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
)

// GetUserByUsername returns a user information associated with id
// HealthCheck godoc
// @Summary Get a user using username.
// @Description returns a user info using username.
// @Tags user
// @Param username path string true "username"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/{username} [get]
func GetUserByUsername(c *fiber.Ctx) error {
	u := new(models.User)
	if err := models.GetUserByUsername(c.Params("username"), u); err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(apiv1.NewUserResponse(u))
}

// UpdateUser updates a user's information
// HealthCheck godoc
// @Summary Update a user
// @Description Update user info
// @Tags user
// @Param username path string true "username"
// @Param data body apiv1.UserUpdateRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/{username} [put]
func UpdateUser(c *fiber.Ctx) error {

	return nil
}

// UpdatePassword changes the password for a given valid user
// HealthCheck godoc
// @Summary Update user password
// @Description updates user info
// @Tags user
// @Param username path string true "username"
// @Param data body apiv1.UserUpdatePasswordRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/{username}/password [put]
func UpdatePassword(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("")
}

// GenerateForgotPasswordLink sends a reset password link to the email
func GenerateForgotPasswordLink(c *fiber.Ctx) error {

	return nil
}

// ResetPasswordFromLink resets the password from a link
func ResetPasswordFromLink(c *fiber.Ctx) error {

	return nil
}
