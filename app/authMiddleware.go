package app

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

type MissingToken struct {
	Message string `json:"message"`
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					utils.WriteJson(w, appError.Code, appError.AsMessage())
				}
			} else {
				utils.WriteJson(w, http.StatusUnauthorized, MissingToken{Message: "missing token"})
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
