package domain

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var HMAC_SAMPLE_SECRET = os.Getenv("HMAC_SAMPLE_SECRET")

type Login struct {
	Email  string `json:"email"db:"email"`
	UserId string `json:"user_id"db:"id"`
	Role   string `json:"role"db:"role"`
}

func (l Login) ClaimsForUser() AccessTokenClaims {
	return AccessTokenClaims{
		UserId: l.UserId,
		Role:   l.Role,
		Email:  l.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenDuration).Unix(),
		},
	}
}

func (l Login) ClaimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		UserId: l.UserId,
		Role:   l.Role,
		Email:  l.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenDuration).Unix(),
		},
	}
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {
	if l.Role == userRole {
		return l.ClaimsForUser()
	} else {
		return l.ClaimsForAdmin()
	}
}
