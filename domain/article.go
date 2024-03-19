package domain

import (
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
)

type Article struct {
	ID          int    `json:"article_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishedAt string `json:"published_at"`
	Author      string `json:"author"`
}

type ArticleRepository interface {
	GetAll() ([]Article, *errs.AppError)
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
