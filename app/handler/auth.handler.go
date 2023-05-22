package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	models2 "github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/pkg/mailer"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
)

// SignUp sign up a user
func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	u, ve := new(models2.User), new(models2.VerifyEmail)
	req := new(userRegisterRequest)
	if err := req.bind(ctx, u, ve, h.validator); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	if err := h.UserStore.GetUserByUsernameOrEmail(u.Username, ve.Email, u); err == nil {
		fmt.Println(u.Username)
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.DuplicateEntry())
	}
	if err := h.UserStore.Create(u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := h.VerifyEmailStore.Create(ve, u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err := mailer.NewMailer().
		To([]string{ve.ExtractEmail()}).
		WithSubject("Codera OJ Email Verification").
		WithTemplate(mailer.EmailTypeEmailVerification, ve).
		Send()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return ctx.Status(http.StatusCreated).JSON("please verify your email to complete the sign up process")
}

// VerifyEmail verifies email
func (h *Handler) VerifyEmail(c *fiber.Ctx) error {
	id, token := c.Params("id"), c.Params("token")
	ve := new(models2.VerifyEmail)
	if err := h.VerifyEmailStore.GetIDToken(id, token, ve); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := ve.IsLinkExpired(); err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}
	u := ve.User
	u.Email = ve.Email
	u.Verified = true
	if err := h.UserStore.UpdateUser(&u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON("email verified successfully")
}

// RefreshToken returns a jwt claims for a valid user
func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	u := new(models2.User)
	req := new(userLoginRequest)
	if err := req.bind(c, u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	if err := h.UserStore.GetUserByUsernameOrEmail(u.Username, u.Email, u); err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}
	if err := u.CheckPassword(req.User.Password); err != nil {
		return c.Status(http.StatusForbidden).JSON(utils.PasswordError())
	}
	encodedToken, err := GenerateEncodedToken(u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"token": encodedToken})
}
