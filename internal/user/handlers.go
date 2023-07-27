package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/utils"
	"github.com/imtiaz246/codera_oj/models"
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

	// todo: change u to APIFormat
	return c.Status(http.StatusOK).JSON(u)
}

// UpdateUser updates a user's information
// HealthCheck godoc
// @Summary Update a user
// @Description Update user info
// @Tags user
// @Param username path string true "username"
// @Param data body UserUpdateRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [put]
func UpdateUser(c *fiber.Ctx) error {

	return nil
}

// UpdatePassword changes the password for a given valid user
// HealthCheck godoc
// @Summary Update user password
// @Description updates user info
// @Tags user
// @Param username path string true "username"
// @Param data body UserUpdatePasswordRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/password [put]
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
