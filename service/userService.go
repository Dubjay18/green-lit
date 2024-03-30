package service

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
)

type UserService interface {
	PopulateUsers() *errs.AppError
	GetAllUsers() ([]dto.UserResponse, *errs.AppError)
	GetByID(id int) (*dto.UserResponse, *errs.AppError)
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

// Todo: Implement pagination
func (s DefaultUserService) GetAllUsers() ([]dto.UserResponse, *errs.AppError) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s DefaultUserService) GetByID(id int) (*dto.UserResponse, *errs.AppError) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errs.NewNotFoundError("User not found")
	}
	return user, nil
}

func NewUserService(repo domain.UserRepositoryDB) DefaultUserService {
	return DefaultUserService{repo}
}
