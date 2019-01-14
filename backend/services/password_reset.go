package services

import (
	"fmt"

	"bitbucket.org/mendelgusmao/me_gu/backend/config"
	gomail "gopkg.in/gomail.v2"
)

const (
	passwordResetEmailTemplate = `
  To reset your password, simply copy and paste the address above into the
	your web browser and follow the instructions.

  %s%s
`
	passwordResetEmailSubject = "Password reset"
)

type PasswordReset struct{}

func (p PasswordReset) SendEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Backend.PasswordResetFromAddress)
	m.SetHeader("To", email)
	m.SetHeader("Subject", passwordResetEmailSubject)
	m.SetBody("text/plain", p.formatMessage(token))

	return gomail.NewDialer(
		config.Backend.SMTPAddress,
		config.Backend.SMTPPort,
		config.Backend.SMTPUser,
		config.Backend.SMTPPassword,
	).DialAndSend(m)
}

func (p PasswordReset) formatMessage(token string) string {
	return fmt.Sprintf(passwordResetEmailTemplate,
		config.Backend.PasswordResetURL,
		token)
}
