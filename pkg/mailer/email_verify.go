package mailer

import (
	"bytes"
	"github.com/imtiaz246/codera_oj/app/models"
	"text/template"
)

// template
// create template
// execute template with data
// return the template

const htmlTemplate = `
<h3>Hello {{ .Username }}.</h3> <br>
Thanks for registering with us! <br>
Please <a href={{ .Link }}>click here</a> to verify your email address before {{ .ExpirationTime }}. <br>
`

type templateData struct {
	Username       string
	Link           string
	ExpirationTime string
}

func initTemplateData(reqData interface{}) templateData {
	ve := reqData.(*models.VerifyEmail)
	data := new(templateData)
	data.Username = ve.User.Username
	data.Link = ve.GenerateLink()
	data.ExpirationTime = ve.ExpirationTime.Format("2006-01-02 03:04:05 pm")
	return *data
}

func createEmailVerifyTemplate(reqData interface{}) ([]byte, error) {
	data := initTemplateData(reqData)
	t, err := template.New("verify-email").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}
	var output bytes.Buffer
	if err := t.Execute(&output, data); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
