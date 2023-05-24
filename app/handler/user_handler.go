package handler

import (
	"github.com/gofiber/fiber/v2"
	apisv1 "github.com/imtiaz246/codera_oj/app/apis/v1"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
)

// GetUserByUsername returns a user information associated with id
func (h *Handler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	u := new(models.User)
	if err := h.UserStore.GetUserByUsername(username, u); err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(apisv1.NewUserResponse(u))
}

// UpdateUser updates a user's information
func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	return nil
}

// UpdatePassword changes the password for a given valid user
func (h *Handler) UpdatePassword(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON("")
}

// GenerateForgotPasswordLink sends a reset password link to the email
func (h *Handler) GenerateForgotPasswordLink(c *fiber.Ctx) error {

	return nil
}

// ResetPasswordFromLink resets the password from a link
func (h *Handler) ResetPasswordFromLink(c *fiber.Ctx) error {

	return nil
}
