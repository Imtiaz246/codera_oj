package mailer

import (
	models2 "github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func init() {
	config.LoadConfigs()
}
func TestMailer(t *testing.T) {
	user := models2.User{
		Username: "imtiaz_email_test",
		Email:    "imtiazuddincho246@gmail.com",
	}
	evm := models2.VerifyEmail{
		Email:          "imtiazuddincho246@gmail.com",
		UserId:         1,
		User:           user,
		ExpirationTime: time.Now().Add(time.Minute * 10),
	}
	evm.GenerateToken()

	err := NewMailer().
		To([]string{user.ExtractEmail()}).
		WithTemplate(EmailTypeEmailVerification, &evm).
		WithSubject("verify email").
		Send()
	require.Equal(t, err, nil)
}
