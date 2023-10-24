package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/adapters/handler/http/utils"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/port"
	"net/http"
)

// AuthHandler represents the HTTP handler for authentication-related requests
type AuthHandler struct {
	svc port.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc,
	}
}

// SignUp signs up a user
// HealthCheck godoc
// @Summary SignUp a user.
// @Description create an account for a user.
// @Tags auth
// @Param data body structs.UserRegisterRequest true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/signup [post]
func (a *AuthHandler) SignUp(ctx *fiber.Ctx) error {
	req := new(dto.UserRegistration)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	return nil
}

// Login create access token and refresh token for a valid user
// HealthCheck godoc
// @Summary Login a user.
// @Description logs in a user if valid credentials given.
// @Tags auth
// @Param data body structs.UserLoginRequest true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} UserLoginResponse
// @Router /auth/login [post]
func (a *AuthHandler) Login(ctx *fiber.Ctx) error {
	req := new(dto.UserLogin)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}

	return nil
}

// VerifyEmail verifies email of a valid user
// HealthCheck godoc
// @Summary Verify email address.
// @Description Verify email address.
// @Tags auth
// @Param id path string true "token ID"
// @Param token path string true "token"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/verify-email/{id}/{token} [get]
func (a *AuthHandler) VerifyEmail(ctx *fiber.Ctx) error {

	return nil
}

// RenewToken renews the access token using a valid refresh token
// HealthCheck godoc
// @Summary Renew the access token
// @Description Renew the access token using the refresh token
// @Tags auth
// @Param username path string true "username"
// @Param refresh-token query string true "refresh token" minlength(1)
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/renew-token [get]
func (a *AuthHandler) RenewToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Query("refresh-token")
	fmt.Print(refreshToken)

	return nil
}
