package domain

import (
	"bytes"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"html/template"
	"net/smtp"
	"sync"
)

type EmailRepositoryDB struct {
}

func (r EmailRepositoryDB) SendEmail(email Email, wg *sync.WaitGroup) *errs.AppError {
	defer wg.Done()

	// Parse and execute the email template
	t, err := template.ParseFiles(email.Template)
	if err != nil {
		return errs.NewUnexpectedError("Unexpected error while parsing email template")
	}
	var body bytes.Buffer
	err = t.Execute(&body, email)
	if err != nil {
		return errs.NewUnexpectedError("Unexpected error while executing email template")
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email.To}
	msg := []byte("To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		body.String() + "\r\n")
	err = smtp.SendMail("smtp.example.com:25", auth, "sender@example.com", to, msg)
	if err != nil {
		logger.Error("Error while sending email: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error while sending email")
	}

	return nil
}

func NewEmailRepositoryDB() EmailRepositoryDB {
	return EmailRepositoryDB{}
}
