package utils

import (
	"fmt"
	"net/smtp"
)

// handles internally all the required logic to send an email
// by given values, it can send html mails
func SendEmail(to []string, topic, message string) error {
	auth := smtp.PlainAuth(
		"",
		GlobalEnv.Email.User,
		GlobalEnv.Email.Password,
		GlobalEnv.Email.Host,
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	body := fmt.Sprintf("Subject: %v\n%v\n\n%v", topic, headers, message)
	address := fmt.Sprintf("%v:%v", GlobalEnv.Email.Host, GlobalEnv.Email.Port)
	err := smtp.SendMail(address, auth, "system@librecode.com", to, []byte(body))

	return err
}
