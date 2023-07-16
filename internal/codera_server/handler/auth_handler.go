package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/custom/config"
	apiv1 "github.com/imtiaz246/codera_oj/internal/codera_server/api/v1"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/modules/mailer"
	"github.com/imtiaz246/codera_oj/modules/token"
	"github.com/imtiaz246/codera_oj/utils"
	"github.com/o1egl/paseto"
	"net/http"
	"time"
)

var (
	ErrTokenIsBlocked  = fmt.Errorf("token is blocked")
	ErrTokenIsNotValid = fmt.Errorf("token is not valid")
)

// SignUp signs up a user
// HealthCheck godoc
// @Summary SignUp a user.
// @Description create account for a user.
// @Tags auth
// @Param data body apiv1.UserRegisterRequest true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/signup [post]
func SignUp(ctx *fiber.Ctx) error {
	req := new(apiv1.UserRegisterRequest)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	u, ve := extractRegistrationRequest(req)
	if err := u.HashPassword(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if err := models.GetUserByUsernameOrEmail(u.Username, ve.Email, u); err == nil {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.DuplicateEntry())
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

	return ctx.Status(http.StatusCreated).JSON(apiv1.UserSuccessfulRegistrationResponse)
}

// Login create access token and refresh token for a valid user
// HealthCheck godoc
// @Summary Login a user.
// @Description logs in a user if valid credentials given.
// @Tags auth
// @Param data body apiv1.UserLoginRequest true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} apiv1.UserLoginResponse
// @Router /auth/login [post]
func Login(ctx *fiber.Ctx) error {
	req := new(apiv1.UserLoginRequest)
	if err := BindAndValidate(ctx, req); err != nil {
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

	return ctx.Status(http.StatusOK).JSON(apiv1.NewLoginResponse(u, accessTokenInfo, refreshTokenInfo))
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

	return ctx.Status(http.StatusOK).JSON(apiv1.EmailSuccessfulVerificationResponse)
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
// @Router /auth/{username}/renew-token [get]
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

	return ctx.Status(http.StatusOK).JSON(apiv1.NewRenewTokenResponse(user, accessTokenInfo))
}

// extractRegistrationRequest extracts information for user registration request
func extractRegistrationRequest(r *apiv1.UserRegisterRequest) (*models.User, *models.VerifyEmail) {
	u := &models.User{
		Username: r.Username,
		Password: r.Password,
	}
	ve := &models.VerifyEmail{
		Email: r.Email,
	}

	return u, ve
}

// sendEmailVerificationMail sends email verification mail to user
func sendEmailVerificationMail(ve *models.VerifyEmail) error {
	return mailer.NewMailer().
		To([]string{ve.ExtractEmail()}).
		WithSubject("Codera OJ Email Verification").
		WithTemplate(mailer.EmailTypeEmailVerification, ve).
		Send()
}

// getTokenManager get the token manager
func getTokenManager(authConfig config.AuthConfig) (token.TokenManager, error) {
	key, err := base64.StdEncoding.DecodeString(authConfig.Key)
	if err != nil {
		return nil, err
	}
	tokenManager, err := token.NewPasetoToken(key)
	if err != nil {
		return nil, err
	}
	return tokenManager, nil
}

// getTokens returns access token and refresh token for a valid user
func getTokens(claimsInfo *token.ClaimsInfo) (accessTokenInfo, refreshTokenInfo *token.TokenInfo, err error) {
	authConfig := config.Settings.Auth
	tokenManager, err := getTokenManager(authConfig)
	if err != nil {
		return
	}

	accessTokenDuration, err := time.ParseDuration(authConfig.AccessTokenDuration)
	if err != nil {
		return
	}
	accessTokenInfo, err = tokenManager.CreateToken(claimsInfo, accessTokenDuration)
	if err != nil {
		return
	}

	refreshTokenDuration, err := time.ParseDuration(authConfig.RefreshTokenDuration)
	if err != nil {
		return
	}
	refreshTokenInfo, err = tokenManager.CreateToken(claimsInfo, refreshTokenDuration)
	if err != nil {
		return
	}

	return
}

// getAccessToken get the access token with claims and returns the TokenInfo
func getAccessToken(claimsInfo *token.ClaimsInfo) (accessTokenInfo *token.TokenInfo, err error) {
	authConfig := config.Settings.Auth
	tokenManager, err := getTokenManager(authConfig)
	if err != nil {
		return
	}

	accessTokenDuration, err := time.ParseDuration(authConfig.AccessTokenDuration)
	if err != nil {
		return
	}
	accessTokenInfo, err = tokenManager.CreateToken(claimsInfo, accessTokenDuration)
	if err != nil {
		return
	}

	return
}

// getTokenPayload verifies the token and returns the paseto json payload
func getPasetoJsonPayload(tokenStr string) (*paseto.JSONToken, error) {
	authConfig := config.Settings.Auth
	key, err := base64.StdEncoding.DecodeString(authConfig.Key)
	if err != nil {
		return nil, err
	}

	tokenManager, err := token.NewPasetoToken(key)
	if err != nil {
		return nil, err
	}

	pasetoPayload, err := tokenManager.VerifyToken(tokenStr)
	if err != nil {
		return nil, err
	}
	return pasetoPayload, nil
}

// createSessionFromTokenInfo creates session from token info
func createSessionFromTokenInfo(tokenInfo *token.TokenInfo) (*models.Session, error) {
	tokenUUID, err := uuid.Parse(tokenInfo.Payload.Jti)
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	if err := models.GetUserByUsername(tokenInfo.Payload.Get("username"), user); err != nil {
		return nil, err
	}
	session := &models.Session{
		ID:        tokenUUID,
		User:      user,
		UserID:    user.ID,
		UserAgent: tokenInfo.Payload.Get("userAgent"),
		ClientIP:  tokenInfo.Payload.Get("clientIP"),
		IsBlocked: false,
		ExpiresAt: tokenInfo.Payload.Expiration,
		CreatedAt: tokenInfo.Payload.IssuedAt,
		UpdatedAt: tokenInfo.Payload.IssuedAt,
	}

	return session, nil
}
