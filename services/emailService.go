package services

import (
	"net/smtp"
)

func emailForVerificationCodeHtml(firstName string, verificationCode string) string {
	return ("<h1>Hi " + firstName + "</h1> <p>,welcome to SleekSpace <span style=\"color: rgb(77, 9, 205);font-size: 16px;font-weight:bold;\">Tube Max</span>. Please enter this code to activate your account: " + verificationCode + "</p></p> Code will expire in 30 minutes. See you soon</p>")
}

func SendVerificationCodeEmail(email string, firstName string, verificationCode string) bool {
	html := emailForVerificationCodeHtml(firstName, verificationCode)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: Confirm your email for Account Registration\n" + headers + "\n\n" + html
	auth := smtp.PlainAuth(
		"",
		"webdevndlovu5@gmail.com",
		"tdjwgzmxbhmgydjl",
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"webdevndlovu5@gmail.com",
		[]string{email},
		[]byte(msg),
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
