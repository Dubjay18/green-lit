package domain

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserId    int    `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
}
