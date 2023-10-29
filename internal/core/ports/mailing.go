package ports

import "github.com/imtiaz246/codera_oj/internal/core/domain/dto"

// MailingAdapter is an interface that provides methods for mailing-related operations.
type MailingAdapter interface {
	// SendEmailVerificationMail sends an email for the purpose of email verification.
	SendEmailVerificationMail(receiverEmail string, data dto.EmailVerificationInfo) error
}
