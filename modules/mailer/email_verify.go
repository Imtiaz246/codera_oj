package mailer

import (
	"bytes"
	"github.com/imtiaz246/codera_oj/models"
	"text/template"
)

const htmlTemplate = `
<h3>Hello {{ .Username }}.</h3> <br>
Thanks for registering with us! <br>
Please <a href={{ .Link }}>click here</a> to verify your email address before {{ .ExpirationTime }}. <br>
`

// templateData holds necessary data for email verification
type templateData struct {
	Username       string
	Link           string
	ExpirationTime string
}

// fillTemplateData fills the templates data
func fillTemplateData(data interface{}) *templateData {
	ve := data.(*models.VerifyEmail)
	d := &templateData{
		Username:       ve.User.Username,
		Link:           ve.GenerateLink(),
		ExpirationTime: ve.ExpirationTime.Format("2006-01-02 03:04:05 pm"),
	}

	return d
}

// createEmailVerifyTemplate creates template with necessary data for verifying email
func createEmailVerifyTemplate(reqData interface{}) ([]byte, error) {
	tmpl, err := template.New("verify-email").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, fillTemplateData(reqData)); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
