package domain

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AccessTokenDuration  = 15
	RefreshTokenDuration = 30
	userRole             = "user"
	adminRole            = "admin"
)

type AccessTokenClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserId    string `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
}

func (c AccessTokenClaims) IsUserRole() bool {
	return c.Role == userRole
}

func (c AccessTokenClaims) IsAdminRole() bool {
	return c.Role == adminRole

}

func (c AccessTokenClaims) isOwner(userId string) bool {
	return c.UserId == userId
}

func (c AccessTokenClaims) isRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	return c.UserId == urlParams["user_id"]
}

func (c AccessTokenClaims) isRequestVerifiedWithTokenClaimsAndRole(urlParams map[string]string) bool {
	return c.UserId == urlParams["user_id"] && c.Role == adminRole
}

func (c AccessTokenClaims) RefreshTokenClaims() RefreshTokenClaims {
	return RefreshTokenClaims{
		UserId:    c.UserId,
		TokenType: "refresh_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(RefreshTokenDuration)).Unix(),
		},
	}
}
