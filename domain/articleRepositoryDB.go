package domain

import (
	"fmt"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	queryGetAllArticles     = "SELECT * FROM articles"
	queryGetArticleByID     = "SELECT * FROM articles WHERE article_id=$1"
	queryGetArticleByUserID = "SELECT * FROM articles WHERE author_id=$1"
	queryCreateArticle      = "INSERT INTO articles (title, content, published_at, author_id) VALUES ($1, $2, $3, $4) RETURNING article_id"
)

type ArticleRepositoryDB struct {
	db *sqlx.DB
}

func (r ArticleRepositoryDB) GetAll() ([]Article, *errs.AppError) {
	var articles []Article
	var err error
	err = r.db.Select(&articles, queryGetAllArticles)
	if err != nil {
		logger.Error("Error while querying articles" + err.Error())
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
	err = r.db.Select(&articles, queryGetArticleByUserID, id)
	if err != nil {
		logger.Error("Error while querying articles" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return articles, nil
}

func (r ArticleRepositoryDB) CreateArticle(article Article) (*Article, *errs.AppError) {
	if article.Title == "" {
		return nil, errs.NewValidationError("Title cannot be empty")
	}
	if article.Content == "" {
		return nil, errs.NewValidationError("Content cannot be empty")
	}
	if article.Author == 0 {
		return nil, errs.NewValidationError("Author cannot be empty")
	}
	var lastId int64
	err := r.db.QueryRow(queryCreateArticle, article.Title, article.Content, article.PublishedAt, article.Author).Scan(&lastId)
	if err != nil {
		logger.Error("Error while creating article" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	if lastId == 0 {
		logger.Error("Error while creating article")
		return nil, errs.NewUnexpectedError("Error while creating article")
	}
	fmt.Println(strconv.FormatInt(lastId, 10), "sjjsjs")
	article.ID = strconv.FormatInt(lastId, 10)
	return &article, nil
}

func NewArticleRepositoryDB(dbClient *sqlx.DB) ArticleRepositoryDB {
	return ArticleRepositoryDB{dbClient}
}
