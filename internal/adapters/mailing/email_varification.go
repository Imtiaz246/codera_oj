package mailing

import (
	"bytes"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"text/template"
)

const emailVerificationTemplate = `
<h3>Hello {{ .Username }}.</h3> <br>
Thanks for registering with us! <br>
Please <a href={{ .VerificationLink }}>click here</a> to verify your email address before {{ .ExpirationTime }}. <br>
`

// SendEmailVerificationMail sends email for email verification using emailVerificationTemplate template
func (m *MailingAdapter) SendEmailVerificationMail(receiverEmail string, data dto.EmailVerificationInfo) error {
	tmpl, err := createEmailVerifyTemplate(data)
	if err != nil {
		return err
	}

	err = m.to([]string{receiverEmail}).
		withSubject("Codera OJ Email Verification").
		withTemplate(tmpl).
		send()

	return err
}

// createEmailVerifyTemplate creates template with necessary data for verifying email
func createEmailVerifyTemplate(data dto.EmailVerificationInfo) ([]byte, error) {
	tmpl, err := template.New("email-verification").Parse(emailVerificationTemplate)
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	if err = tmpl.Execute(&output, data); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
