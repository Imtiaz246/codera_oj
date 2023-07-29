package mailer

import (
	"github.com/imtiaz246/codera_oj/models"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMailer(t *testing.T) {
	user := models.User{
		Handle: "imtiaz_email_test",
		Email:  "imtiazuddincho246@gmail.com",
	}
	evm := models.VerifyEmail{
		Email:          "imtiazuddincho246@gmail.com",
		UserId:         1,
		User:           user,
		ExpirationTime: time.Now().Add(time.Minute * 10),
	}
	err := evm.GenerateToken()
	require.Equal(t, err, nil)

	err = NewMailer().
		To([]string{user.ExtractEmail()}).
		WithTemplate(EmailTypeEmailVerification, &evm).
		WithSubject("verify email").
		Send()
	require.Equal(t, err, nil)
}
