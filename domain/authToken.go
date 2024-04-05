package domain

import (
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	token *jwt.Token
}

func (a AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := a.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("Error signing token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}