package ports

import "github.com/imtiaz246/codera_oj/internal/core/domain/dto"

// MailingAdapter is an adapter to do the mailing-related work
type MailingAdapter interface {
	// SendEmailVerificationMail sends email for email verification
	SendEmailVerificationMail(receiverEmail string, data dto.EmailVerificationInfo) error
}
