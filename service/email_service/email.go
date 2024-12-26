package email_service

import (
	"errors"
	"fmt"
	"net/smtp"
)


func SendEmail(toEmail string, fromEmail string, passwordEmail string, subjectEmail string, bodyEmail string) error {
	from := fromEmail
	password := passwordEmail

	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port


	message := "From: " + from + "\r\n" +
		"To: " + toEmail + "\r\n" +
		"Subject: " + subjectEmail + "\r\n\r\n" +
		bodyEmail

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return errors.New("erro ao enviar email")
	}

	return nil
}

