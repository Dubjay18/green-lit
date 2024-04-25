package service

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/errs"
	"sync"
)

type EmailService interface {
	SendEmail(email domain.Email) *errs.AppError
}

type DefaultEmailService struct {
	repo domain.EmailRepository
}

func (s DefaultEmailService) SendEmail(email domain.Email) *errs.AppError {
	var wg sync.WaitGroup
	wg.Add(1)

	var err *errs.AppError

	go func() {
		defer wg.Done()
		err = s.repo.SendEmail(email, &wg)
	}()

	wg.Wait()

	return err
}

func NewEmailService(repo domain.EmailRepository) DefaultEmailService {
	return DefaultEmailService{repo}
}
