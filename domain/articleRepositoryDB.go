package domain

import (
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetAllArticles     = "SELECT * FROM articles"
	queryGetArticleByID     = "SELECT * FROM articles WHERE id=$1"
	queryGetArticleByUserID = "SELECT * FROM articles WHERE author_id=$1"
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

func (r ArticleRepositoryDB) GetByID(id int) (*dto.ArticleResponse, *errs.AppError) {
	var article dto.ArticleResponse
	err := r.db.Get(&article, queryGetArticleByID, id)
	if err != nil {
		logger.Error("Error while querying article" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &article, nil

}

func (r ArticleRepositoryDB) GetByUserID(id int) ([]Article, *errs.AppError) {
	var articles []Article
	var err error
	err = r.db.Select(&articles, queryGetUserByID, id)
	if err != nil {
		logger.Error("Error while querying articles" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return articles, nil
}

func NewArticleRepositoryDB(dbClient *sqlx.DB) ArticleRepositoryDB {
	return ArticleRepositoryDB{dbClient}
}
