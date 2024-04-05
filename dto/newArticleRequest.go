package dto

import "github.com/Dubjay18/green-lit/errs"

type NewArticleRequest struct {
	Title   string `json:"title"db:"title""`
	Content string `json:"content"db:"content"`
	Author  int    `json:"author"db:"author_id"`
}

func (a NewArticleRequest) Validate() *errs.AppError {
	if a.Title == "" {
		return errs.NewValidationError("Title cannot be empty")
	}
	if a.Content == "" {
		return errs.NewValidationError("Content cannot be empty")
	}
	if a.Author == 0 {
		return errs.NewValidationError("Author cannot be empty")
	}
	return nil
}
