package domain

import (
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetAllArticles = "SELECT * FROM articles"
)

type ArticleRepositoryDB struct {
	db *sqlx.DB
}

func (r ArticleRepositoryDB) GetAll() ([]Article, *errs.AppError) {
	var articles []Article
	var err error
	err = r.db.Select(&articles, queryGetAllArticles)
	if err != nil {
		logger.Error("Error while quering articles" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return articles, nil
}

func NewArticleRepositoryDB(dbClient *sqlx.DB) ArticleRepositoryDB {
	return ArticleRepositoryDB{dbClient}
}
