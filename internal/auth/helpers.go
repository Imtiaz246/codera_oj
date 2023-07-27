package auth

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/custom/config"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/modules/mailer"
	"github.com/imtiaz246/codera_oj/modules/token"
	"github.com/o1egl/paseto"
	"time"
)

// extractRegistrationRequest extracts information for user registration request
func extractRegistrationRequest(r *UserRegisterRequest) (*models.User, *models.VerifyEmail) {
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
