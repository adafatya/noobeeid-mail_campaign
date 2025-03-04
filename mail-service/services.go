package main

import "gopkg.in/gomail.v2"

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_AUTH_EMAIL    = "aufaditya1895@gmail.com"
	CONFIG_AUTH_PASSWORD = "APP_PASSWORD"
)

func sendMailGoMail(from string, to []string, subject string, message string, bodyType string) (err error) {
	// setup gomail message
	mailer := gomail.NewMessage()
	// setting header from
	mailer.SetHeader("From", from)
	// setting header to
	mailer.SetHeader("To", to...)

	// setting subject
	mailer.SetHeader("Subject", subject)

	// setting body
	var mimeBodyType string
	if bodyType == "html" {
		mimeBodyType = "text/html"
	} else {
		mimeBodyType = "text/plain"
	}

	mailer.SetBody(mimeBodyType, message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	return
}
