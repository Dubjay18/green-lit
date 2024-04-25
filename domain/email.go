package domain

import (
	"github.com/Dubjay18/green-lit/errs"
	"sync"
)

type Email struct {
	To       string
	From     string
	Subject  string
	Body     string
	Template string
}

type EmailRepository interface {
	SendEmail(email Email, wg *sync.WaitGroup) *errs.AppError
}
