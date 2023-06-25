package mailer

import (
	"github.com/imtiaz246/codera_oj/models"
	"strings"
)

func getEmailDomain(emailAddr string) string {
	return strings.Split(emailAddr, "@")[1]
}

func extractEmailAddr(us []*models.User) []string {
	var ea []string
	for _, u := range us {
		ea = append(ea, u.Email)
	}
	return ea
}
