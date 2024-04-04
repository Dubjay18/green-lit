package domain

import (
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
)

type Article struct {
	ID          string `json:"article_id"db:"article_id"`
	Title       string `json:"title"db:"title""`
	Content     string `json:"content"db:"content"`
	PublishedAt string `json:"published_at"db:"published_at"`
	Author      int    `json:"author"db:"author_id"`
}

type ArticleRequest struct {
	Title   string `json:"title"db:"title""`
	Content string `json:"content"db:"content"`
	Author  int    `json:"author"db:"author_id"`
}

type ArticleRepository interface {
	GetAll() ([]Article, *errs.AppError)
	GetByID(id int) (*Article, *errs.AppError)
	GetByUserID(id int) ([]Article, *errs.AppError)
	CreateArticle(article ArticleRequest) *errs.AppError
}

func (a Article) ToDto() dto.ArticleResponse {
	return dto.ArticleResponse{
		ID:          a.ID,
		Title:       a.Title,
		Content:     a.Content,
		PublishedAt: a.PublishedAt,
		Author:      a.Author,
	}

}
