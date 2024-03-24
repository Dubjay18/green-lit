package service

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/errs"
)

type UserService interface {
	PopulateUsers() *errs.AppError
}

type DefaultUserService struct {
	repo domain.UserRepositoryDB
}

func (s DefaultUserService) PopulateUsers() *errs.AppError {
	err := s.repo.Populate()
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(repo domain.UserRepositoryDB) DefaultUserService {
	return DefaultUserService{repo}
}
