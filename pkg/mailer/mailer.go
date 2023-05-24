package mailer

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	gmailSMTPAuth              = "smtp.gmail.com"
	gmailSMTPServer            = "smtp.gmail.com:587"
	emailSourceName            = "codera OJ"
	EmailTypeContestReminder   = "ContestReminder"
	EmailTypeEmailVerification = "EmailVerification"
)

type Mailer interface {
	Send() error
	To([]string) *mail
	WithSubject(string) *mail
	WithTemplate(string, interface{}) *mail
	WithAttachments([]string) *mail
}

type mail struct {
	senderName string
	senderAddr string
	senderPass string
	error      error
	email      email.Email
}

func NewMailer() Mailer {
	emailConfig := config.GetEmailConfig()
	return &mail{
		senderName: emailSourceName,
		senderAddr: emailConfig.SenderAddr,
		senderPass: emailConfig.SenderPass,
		error:      nil,
	}
}

func (m *mail) Send() error {
	if m.error != nil {
		return m.error
	}
	m.email.From = fmt.Sprintf("%s <%s>", m.senderName, m.senderAddr)
	smtpAuth := smtp.PlainAuth("", m.senderAddr, m.senderPass, gmailSMTPAuth)
	return m.email.Send(gmailSMTPServer, smtpAuth)
}

func (m *mail) To(es []string) *mail {
	m.email.To = es
	return m
}

func (m *mail) WithCC(us []*models.User) *mail {
	cc := extractEmailAddr(us)
	m.email.Cc = cc
	return m
}

func (m *mail) WithBCC(us []*models.User) *mail {
	bcc := extractEmailAddr(us)
	m.email.Bcc = bcc
	return m
}

func (m *mail) WithSubject(sub string) *mail {
	m.email.Subject = sub
	return m
}

func (m *mail) WithTemplate(templateType string, data interface{}) *mail {
	if m.error != nil {
		return m
	}
	switch templateType {
	case EmailTypeEmailVerification:
		template, err := createEmailVerifyTemplate(data)
		if err != nil {
			m.error = err
			goto END
		}
		m.email.HTML = template
	}
END:
	return m
}

func (m *mail) WithAttachments(files []string) *mail {
	if m.error != nil {
		return m
	}
	for _, f := range files {
		_, err := m.email.AttachFile(f)
		if err != nil {
			m.error = err
			goto END
		}
	}
END:
	return m
}
