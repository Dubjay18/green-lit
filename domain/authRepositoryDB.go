package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/jmoiron/sqlx"
)

const (
	sqlVerifyQry = "SELECT email, password,id FROM users WHERE email = $1"
	sqlSelect    = "select refresh_token from refresh_token_store where refresh_token = $1"
)

type LoginDb struct {
	Email    string `json:"email"db:"email"`
	UserId   string `json:"user_id"db:"id"`
	Role     string `json:"role"db:"role"`
	Password string `db:"password"`
}

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
	var dbLogin LoginDb
	dbLogin.Role = "admin"
	hp, _ := utils.HashPassword(password)
	fmt.Println(hp)
	fmt.Println(email)
	err := a.db.Get(&dbLogin, sqlVerifyQry, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(err.Error())
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	match := utils.CheckPasswordHash(password, dbLogin.Password)
	fmt.Println(dbLogin, match)
	if !utils.CheckPasswordHash(password, dbLogin.Password) {
		return nil, errs.NewAuthenticationError("invalid credentialsde")

	}
	login.Email = dbLogin.Email
	login.Role = dbLogin.Role
	login.UserId = dbLogin.UserId
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
	sqlInsert := sqlSelect
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
