package mailer

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/custom/config"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	GmailSMTPAuth              = "smtp.gmail.com"
	GmailSMTPServer            = "smtp.gmail.com:587"
	EmailSourceName            = "codera OJ"
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

func NewMailer() Mailer {
	emailConfig := config.Settings.Email
	return &mail{
		senderName: EmailSourceName,
		senderAddr: emailConfig.SenderEmail,
		senderPass: emailConfig.SenderPass,
		error:      nil,
	}
}

type mail struct {
	senderName string
	senderAddr string
	senderPass string
	error      error
	email      email.Email
}

func (m *mail) Send() error {
	if m.error != nil {
		return m.error
	}
	smtpAuth := smtp.PlainAuth("", m.senderAddr, m.senderPass, GmailSMTPAuth)
	m.email.From = fmt.Sprintf("%s <%s>", m.senderName, m.senderAddr)

	return m.email.Send(GmailSMTPServer, smtpAuth)
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
