package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera/structs"
	"github.com/imtiaz246/codera_oj/internal/utils"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/modules/token"
	"net/http"
)

var (
	ErrTokenIsBlocked  = fmt.Errorf("token is blocked")
	ErrTokenIsNotValid = fmt.Errorf("token is not valid")
)

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
func SignUp(ctx *fiber.Ctx) error {
	req := new(structs.UserRegisterRequest)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	u, ve := extractRegistrationRequest(req)
	if err := u.HashPassword(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := models.GetUserByUsernameOrEmail(u.Username, ve.Email, u); err == nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}
	if err := models.CreateRecord[*models.User](u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := ve.FillEmailVerifierInfo(u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := models.CreateRecord[*models.VerifyEmail](ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := sendEmailVerificationMail(ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(structs.UserSuccessfulRegistrationResponse)
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
func Login(ctx *fiber.Ctx) error {
	req := new(structs.UserLoginRequest)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}

	u := new(models.User)
	if err := models.GetUserByUsernameOrEmail(req.Username, req.Email, u); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	if err := u.CheckPassword(req.Password); err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(utils.NewError(err))
	}

	newClaimsInfo := &token.ClaimsInfo{
		Username:  u.Username,
		ClientIP:  ctx.IP(),
		UserAgent: ctx.GetReqHeaders()["User-Agent"],
	}
	accessTokenInfo, refreshTokenInfo, err := getTokens(newClaimsInfo)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	session, err := createSessionFromTokenInfo(refreshTokenInfo)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := models.SessionCache.Set(session.ID.String(), *session); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := models.CreateRecord[*models.Session](session); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(structs.NewLoginResponse(u, accessTokenInfo, refreshTokenInfo))
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
func VerifyEmail(ctx *fiber.Ctx) error {
	ve := new(models.VerifyEmail)

	if err := models.GetVerifyEmailUsingIDToken(ctx.Params("id"), ctx.Params("token"), ve); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := ve.IsLinkExpired(); err != nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(err))
	}
	u := ve.VerifiedUser()
	if err := models.UpdateRecord[*models.User](u); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(structs.EmailSuccessfulVerificationResponse)
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
func RenewToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Query("refresh-token")
	pasetoPayload, err := getPasetoJsonPayload(refreshToken)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	claimsInfo := &token.ClaimsInfo{
		Username:  pasetoPayload.Get("username"),
		ClientIP:  pasetoPayload.Get("clientIP"),
		UserAgent: pasetoPayload.Get("userAgent"),
	}
	tokenID := pasetoPayload.Jti

	// Get the session from session_cache or database
	session, err := models.SessionCache.Get(tokenID)
	if err != nil {
		session, err = models.GetRecordByID[*models.Session](tokenID)
		if err != nil {
			return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(ErrTokenIsNotValid))
		}

		if err := models.SessionCache.Set(tokenID, *session); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
		}
	}
	if session.IsBlocked {
		return ctx.Status(http.StatusUnauthorized).JSON(utils.NewError(ErrTokenIsBlocked))
	}

	// Check for token corruption
	if ctx.IP() != claimsInfo.ClientIP || ctx.GetReqHeaders()["User-Agent"] != claimsInfo.UserAgent {
		session.IsBlocked = true
		if err := models.SessionCache.Set(session.ID.String(), *session); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
		}
		if err := models.UpdateRecord[*models.Session](session); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
		}
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(ErrTokenIsBlocked))
	}

	accessTokenInfo, err := getAccessToken(claimsInfo)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	user := new(models.User)
	if err := models.GetUserByUsername(claimsInfo.Username, user); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(structs.NewRenewTokenResponse(user, accessTokenInfo))
}
