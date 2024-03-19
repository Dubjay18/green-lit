package service

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
)

type ArticleService interface {
	GetAllArticles() ([]dto.ArticleResponse, *errs.AppError)
}

type DefaultArticleService struct {
	repo domain.ArticleRepositoryDB
}

func (s DefaultArticleService) GetAllArticles() ([]dto.ArticleResponse, *errs.AppError) {
	articles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.ArticleResponse, 0)
	for _, a := range articles {
		response = append(response, a.ToDto())
	}
	return response, nil

}

func NewArticleService(repo domain.ArticleRepositoryDB) DefaultArticleService {
	return DefaultArticleService{repo}
}
