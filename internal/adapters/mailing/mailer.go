package mailing

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
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

type MailingAdapter struct {
	senderName   string
	serviceEmail string
	servicePass  string
	error        error
	email        email.Email
}

var _ ports.MailingAdapter = (*MailingAdapter)(nil)

func NewMailer(serviceEmail, servicePass string) *MailingAdapter {
	return &MailingAdapter{
		senderName:   EmailSourceName,
		serviceEmail: serviceEmail,
		servicePass:  servicePass,
		error:        nil,
	}
}

func (m *MailingAdapter) send() error {
	if m.error != nil {
		return m.error
	}
	smtpAuth := smtp.PlainAuth("", m.serviceEmail, m.servicePass, GmailSMTPAuth)
	m.email.From = fmt.Sprintf("%s <%s>", m.senderName, m.serviceEmail)

	return m.email.Send(GmailSMTPServer, smtpAuth)
}

func (m *MailingAdapter) to(es []string) *MailingAdapter {
	m.email.To = es
	return m
}

func (m *MailingAdapter) withCC(ccs []string) *MailingAdapter {
	m.email.Cc = ccs
	return m
}

func (m *MailingAdapter) withBCC(bccs []string) *MailingAdapter {
	m.email.Bcc = bccs
	return m
}

func (m *MailingAdapter) withSubject(sub string) *MailingAdapter {
	m.email.Subject = sub
	return m
}

func (m *MailingAdapter) withTemplate(template []byte) *MailingAdapter {
	m.email.HTML = template

	return m
}

func (m *MailingAdapter) withAttachments(files []string) *MailingAdapter {
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
