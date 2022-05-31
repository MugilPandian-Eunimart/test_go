package helpers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func MailSender() error {

	// Sender data.
	from := "testhrmail1@gmail.com"
	password := "@Scientistcr7"

	// Receiver email address.
	to := []string{
		"manisha.s@eunimart.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Important Please Revert Back \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Guess Singh",
		Message: "Completed Work Ah",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return err
}