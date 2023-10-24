package auth

import (
	"context"
	"github.com/imtiaz246/codera_oj/internal/adapters/config"
	"github.com/imtiaz246/codera_oj/internal/adapters/mailing"
	"github.com/imtiaz246/codera_oj/internal/adapters/repo"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/port"
)

type AuthService struct {
	userRepo        repo.UserRepoInterface
	verifyEmailRepo repo.VerifyEmailRepoInterface
	mailingAdptr    port.MailingAdapter
	configAdptr     port.ConfigAdapter
}

var _ port.AuthService = (*AuthService)(nil)

func NewAuthService(ur repo.UserRepoInterface,
	vr repo.VerifyEmailRepoInterface,
	ma *mailing.MailingAdapter,
	ca *config.ConfigAdapter) port.AuthService {
	return &AuthService{
		userRepo:        ur,
		verifyEmailRepo: vr,
		mailingAdptr:    ma,
		configAdptr:     ca,
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
	// TODO: handle for token exists case
	if err = ve.SetVerificationToken(); err != nil {
		return err
	}
	if err = as.verifyEmailRepo.CreateRecord(ve); err != nil {
		return err
	}
	// TODO: send parameter using function by getting dto
	if err = as.mailingAdptr.SendEmailVerificationMail(ve.Email, dto.EmailVerificationInfo{UserName: user.Username}); err != nil {
		return err
	}

	return nil
}

func (as *AuthService) Login(ctx context.Context, data dto.UserLogin) (*dto.UserLoginResponse, error) {

	return nil, nil
}

func (as *AuthService) VerifyEmail(ctx context.Context, id int64, token string) error {

	return nil
}

func (as *AuthService) RenewToken(ctx context.Context, refreshToken string) (string, error) {

	return "", nil
}
