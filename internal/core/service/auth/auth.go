package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/imtiaz246/codera_oj/internal/adapters/mailing"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
)

type AuthService struct {
	userRepo        ports.UserRepoInterface
	verifyEmailRepo ports.VerifyEmailRepoInterface
	mailingAdptr    ports.MailingAdapter
	tokenAdptr      ports.TokenAdapter
	sessionRepo     ports.SessionRepoInterface
}

var _ ports.AuthService = (*AuthService)(nil)

func NewAuthService(ur ports.UserRepoInterface,
	vr ports.VerifyEmailRepoInterface,
	ma *mailing.MailingAdapter,
	ta ports.TokenAdapter,
	sr ports.SessionRepoInterface) ports.AuthService {
	return &AuthService{
		userRepo:        ur,
		verifyEmailRepo: vr,
		mailingAdptr:    ma,
		tokenAdptr:      ta,
		sessionRepo:     sr,
	}
}

func (as *AuthService) SignUp(ctx context.Context, data *dto.UserRegistration) error {
	u := &models.User{
		Username: data.Username,
		Password: data.Password,
	}
	if err := u.HashPassword(); err != nil {
		return err
	}
	user, err := as.userRepo.GetUserByUsernameOrEmail(data.Username, data.Password)
	if err != nil {
		return err
	}
	if err = as.userRepo.CreateRecord(user); err != nil {
		return err
	}

	ve := &models.VerifyEmail{
		Email: data.Email,
		User:  *user,
	}
	ve.SetExpirationTime()
	if err = ve.SetVerificationToken(); err != nil {
		return err
	}
	if err = as.verifyEmailRepo.CreateRecord(ve); err != nil {
		return err
	}
	err = as.mailingAdptr.SendEmailVerificationMail(ve.Email, dto.EmailVerificationInfo{UserName: user.Username})
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) Login(ctx context.Context, data dto.UserLogin) (*dto.UserLoginResponse, error) {
	user, err := as.userRepo.GetUserByUsernameOrEmail(data.Username, data.Email)
	if err != nil {
		return nil, err
	}
	if err = user.CheckPassword(data.Password); err != nil {
		return nil, err
	}
	tokenClaims := &dto.TokenClaims{
		Username:  data.Username,
		ClientIP:  data.ClientIP,
		UserAgent: data.UserAgent,
	}

	accessTokenInfo, err := as.tokenAdptr.CreateAccessToken(tokenClaims)
	if err != nil {
		return nil, err
	}
	refreshTokenInfo, err := as.tokenAdptr.CreateRefreshToken(tokenClaims)
	if err != nil {
		return nil, err
	}
	session := &models.Session{
		ID:        refreshTokenInfo.ID,
		UserID:    user.ID,
		UserAgent: refreshTokenInfo.UserAgent,
		ClientIP:  refreshTokenInfo.ClientIP,
		IsBlocked: false,
		ExpiresAt: refreshTokenInfo.Expiration,
		CreatedAt: refreshTokenInfo.IssuedAt,
		UpdatedAt: refreshTokenInfo.IssuedAt,
	}
	if err = as.sessionRepo.CreateRecord(session); err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{
		User:             user.ToAPIFormat(),
		AccessTokenInfo:  *accessTokenInfo,
		RefreshTokenInfo: *refreshTokenInfo,
	}, nil
}

func (as *AuthService) VerifyEmail(ctx context.Context, id int64, token string) error {
	ve, err := as.verifyEmailRepo.GetVerifyEmailRecordUsingIdToken(id, token)
	if err != nil {
		return fmt.Errorf("verify email record not found: `%v`", err)
	}
	if ve.IsLinkExpired() {
		return fmt.Errorf("link is expired")
	}
	user := &ve.User
	user.Email = ve.Email
	user.Verified = true

	if err = as.userRepo.UpdateRecord(user); err != nil {
		return err
	}

	return nil
}

func (as *AuthService) RenewToken(ctx context.Context, refToken, reqIP, reqUserAgent string) (*dto.TokenInfo, error) {
	tokenInfo, err := as.tokenAdptr.VerifyToken(refToken)
	if err != nil {
		return nil, err
	}
	session, err := as.sessionRepo.GetSessionByTokenUUID(tokenInfo.ID)
	if err != nil {
		return nil, err
	}
	if session.IsBlocked {
		return nil, errors.New("token is blocked")
	}
	// Check for token corruption
	if tokenInfo.ClientIP != reqIP || tokenInfo.UserAgent != reqUserAgent {
		session.IsBlocked = true
		if err = as.sessionRepo.UpdateRecord(session); err != nil {
			return nil, err
		}
		return nil, errors.New("token is blocked")
	}

	claims := &dto.TokenClaims{
		Username:  tokenInfo.Username,
		ClientIP:  tokenInfo.ClientIP,
		UserAgent: tokenInfo.UserAgent,
	}
	accessTokenInfo, err := as.tokenAdptr.CreateAccessToken(claims)
	if err != nil {
		return nil, err
	}

	return accessTokenInfo, nil
}
