package domain

import (
	"database/sql"
	"errors"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/jmoiron/sqlx"
)

const (
	sqlVerifyQry = "SELECT email, password FROM users WHERE email = $1 and password = $1"
	sqlSelect    = "select refresh_token from refresh_token_store where refresh_token = $1"
)

type AuthRepository interface {
	FindById(email string, password string) (*Login, *errs.AppError)
	GenerateAndSaveRefreshTokenToStore(token AuthToken) (string, *errs.AppError)
	RefreshTokenExists(refreshToken string) *errs.AppError
}
type AuthRepositoryDB struct {
	db *sqlx.DB
}

func (a AuthRepositoryDB) FindById(email string, password string) (*Login, *errs.AppError) {
	var login Login

	err := a.db.Get(&login, sqlVerifyQry, email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &login, nil

}

func (a AuthRepositoryDB) RefreshTokenExists(refreshToken string) *errs.AppError {
	sqlSelect := "select refresh_token from refresh_token_store where refresh_token = $1"
	var token string
	err := a.db.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewAuthenticationError("refresh token not registered in the store")
		} else {
			logger.Error("Unexpected database error: " + err.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (a AuthRepositoryDB) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError) {
	// generate the refresh token
	var appErr *errs.AppError
	var refreshToken string
	if refreshToken, appErr = authToken.NewRefreshToken(); appErr != nil {
		return "", appErr
	}

	// store it in the store
	sqlInsert := "insert into refresh_token_store (refresh_token) values ($1)"
	_, err := a.db.Exec(sqlInsert, refreshToken)
	if err != nil {
		logger.Error("unexpected database error: " + err.Error())
		return "", errs.NewUnexpectedError("unexpected database error")
	}
	return refreshToken, nil
}
func NewAuthRepositoryDB(client *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client}
}
