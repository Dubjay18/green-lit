package app

import (
	"fmt"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/logger"
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := utils.DecodeJson(r, &loginRequest)
	if err != nil {
		logger.Error("Error while decoding login request: " + err.Message)
		utils.WriteJson(w, err.Code, err.AsMessage())

	} else {
		token, appErr := h.service.Login(loginRequest)
		if appErr != nil {
			utils.WriteJson(w, appErr.Code, appErr.AsMessage())
		} else {
			utils.WriteJson(w, http.StatusOK, *token)
		}
	}
}

func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	// converting from Query to map type
	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		appErr := h.service.Verify(urlParams)
		fmt.Println(appErr, urlParams)
		if appErr != nil {
			utils.WriteJson(w, appErr.Code, notAuthorizedResponse(appErr.Message))
		} else {
			utils.WriteJson(w, http.StatusOK, authorizedResponse())
		}
	} else {
		utils.WriteJson(w, http.StatusForbidden, notAuthorizedResponse("missing token"))
	}
}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}
