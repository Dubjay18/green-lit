package domain

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/jmoiron/sqlx"
	"net/http"
	"net/url"
)

const (
	sqlVerifyQry = "SELECT email, password,id FROM users WHERE email = $1"
	sqlSelect    = "select refresh_token from refresh_token_store where refresh_token = $1"
)

type LoginDb struct {
	Email    string `json:"email"db:"email"`
	UserId   string `json:"user_id"db:"id"`
	Role     string `json:"role"db:"role"`
	Password string `json:"password"db:"password"`
}

type AuthRepository interface {
	FindById(email string, password string) (*Login, *errs.AppError)
	GenerateAndSaveRefreshTokenToStore(token AuthToken) (string, *errs.AppError)
	RefreshTokenExists(refreshToken string) *errs.AppError
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}
type AuthRepositoryDB struct {
	db *sqlx.DB
}

func (a AuthRepositoryDB) FindById(email string, password string) (*Login, *errs.AppError) {
	var login Login
	var dbLogin LoginDb
	dbLogin.Role = "admin"
	p, _ := utils.HashPassword("password")
	hp, _ := utils.HashPassword(password)
	fmt.Println(hp)
	fmt.Println(utils.CheckPasswordHash("password", p))
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
	fmt.Println(dbLogin, match, password)
	if !match {
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
func (r AuthRepositoryDB) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			fmt.Println(response)
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

/*
This will generate a url for token verification in the below format

/auth/verify?token={token string}

	&routeName={current route name}
	&customer_id={customer id from the current route}
	&account_id={account id from current route if available}

Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8080", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	return u.String()
}
func NewAuthRepositoryDB(client *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client}
}
