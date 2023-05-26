package handler

import (
	"github.com/gofiber/fiber/v2"
	apisv1 "github.com/imtiaz246/codera_oj/app/apis/v1"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/pkg/mailer"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
)

// SignUp sign up a user
func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	r := new(apisv1.UserRegisterRequest)
	if err := BindAndValidate(ctx, r); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	u, ve := extractRegisterRequest(r)
	if err := h.UserStore.GetUserByUsernameOrEmail(u.Username, ve.Email, u); err == nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.DuplicateEntry())
	}
	if err := h.UserStore.Create(u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := ve.FillEmailVerifierInfo(u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := h.VerifyEmailStore.Create(ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := sendVerificationMail(ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(apisv1.UserSuccessfulRegistrationMessage)
}

// VerifyEmail verifies email
func (h *Handler) VerifyEmail(c *fiber.Ctx) error {
	ve := new(models.VerifyEmail)
	if err := h.VerifyEmailStore.GetIDToken(c.Params("id"), c.Params("token"), ve); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := ve.IsLinkExpired(); err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}
	u := ve.VerifiedUser()
	if err := h.UserStore.UpdateUser(u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(apisv1.EmailSuccessfulVerificationMessage)
}

func extractRegisterRequest(r *apisv1.UserRegisterRequest) (*models.User, *models.VerifyEmail) {
	u := new(models.User)
	ve := new(models.VerifyEmail)
	u.Username = r.User.Username
	u.Password = r.User.Password
	ve.Email = r.User.Email

	return u, ve
}

func sendVerificationMail(ve *models.VerifyEmail) error {
	return mailer.NewMailer().
		To([]string{ve.ExtractEmail()}).
		WithSubject("Codera OJ Email Verification").
		WithTemplate(mailer.EmailTypeEmailVerification, ve).
		Send()
}
