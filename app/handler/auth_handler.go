package handler

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	apisv1 "github.com/imtiaz246/codera_oj/app/apis/v1"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/imtiaz246/codera_oj/services/mailer"
	"github.com/imtiaz246/codera_oj/services/token"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
	"time"
)

// SignUp signs up a user
// HealthCheck godoc
// @Summary SignUp a user.
// @Description create account for a user.
// @Tags auth
// @Param data body apisv1.UserRegisterRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/signup [post]
func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	req := new(apisv1.UserRegisterRequest)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	u, ve := extractRegistrationRequest(req)
	if err := u.HashPassword(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
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
	if err := sendEmailVerificationMail(ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(apisv1.UserSuccessfulRegistrationResponse)
}

// Login create access token and refresh token for a valid user
// HealthCheck godoc
// @Summary Login a user.
// @Description logs in a user if valid credentials given.
// @Tags auth
// @Param data body apisv1.UserLoginRequest true "data"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *Handler) Login(ctx *fiber.Ctx) error {
	req := new(apisv1.UserLoginRequest)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}

	u := new(models.User)
	if err := h.UserStore.GetUserByUsernameOrEmail(req.Username, req.Email, u); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	if err := u.CheckPassword(req.Password); err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(utils.NewError(err))
	}

	// todo: add sessions and cookies
	at, rt, err := getTokens(u)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return ctx.Status(http.StatusOK).JSON(apisv1.NewLoginResponse(u, at, rt))
}

// VerifyEmail verifies email of a valid user
// HealthCheck godoc
// @Summary Verify email address.
// @Description Verify email address.
// @Tags auth
// @Param id path string true "token ID"
// @Param token path string true "token"
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/verify-email/{id}/{token} [get]
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

	return c.Status(http.StatusOK).JSON(apisv1.EmailSuccessfulVerificationResponse)
}

func extractRegistrationRequest(r *apisv1.UserRegisterRequest) (*models.User, *models.VerifyEmail) {
	u := &models.User{
		Username: r.Username,
		Password: r.Password,
	}
	ve := &models.VerifyEmail{
		Email: r.Email,
	}

	return u, ve
}

func sendEmailVerificationMail(ve *models.VerifyEmail) error {
	return mailer.NewMailer().
		To([]string{ve.ExtractEmail()}).
		WithSubject("Codera OJ Email Verification").
		WithTemplate(mailer.EmailTypeEmailVerification, ve).
		Send()
}

func getTokens(u *models.User) (at *token.TokenInfo, rt *token.TokenInfo, err error) {
	cfg := config.GetAuthConfig()
	key, err := base64.StdEncoding.DecodeString(cfg.Key)
	if err != nil {
		return
	}
	tm, err := token.NewPasetoToken(key)
	if err != nil {
		return
	}

	ad, err := time.ParseDuration(cfg.AccessTokenDuration)
	if err != nil {
		return
	}
	at, err = tm.CreateToken(u.Username, ad)
	if err != nil {
		return
	}

	rd, err := time.ParseDuration(cfg.RefreshTokenDuration)
	if err != nil {
		return
	}
	rt, err = tm.CreateToken(u.Username, rd)
	if err != nil {
		return
	}

	return
}
