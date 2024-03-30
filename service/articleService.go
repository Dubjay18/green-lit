package service

import (
	"fmt"
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
)

type ArticleService interface {
	GetAllArticles() ([]dto.ArticleResponse, *errs.AppError)
	GetByID(id int) (*dto.ArticleResponse, *errs.AppError)
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

func (s DefaultArticleService) GetByID(id int) (*dto.ArticleResponse, *errs.AppError) {
	article, err := s.repo.GetByID(id)
	if err != nil {
		fmt.Sprintf("Error while getting article: %s", err.Message)
		return nil, err
	}
	if article == nil {
		return nil, errs.NewNotFoundError("Article not found")
	}
	return article, nil
}

func NewArticleService(repo domain.ArticleRepositoryDB) DefaultArticleService {
	return DefaultArticleService{repo}
}
