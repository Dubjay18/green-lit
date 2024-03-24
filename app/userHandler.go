package app

import (
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func (h UserHandler) PopulateUsers(w http.ResponseWriter, r *http.Request) {
	err := h.service.PopulateUsers()
	if err != nil {
		utils.WriteJson(w, err.Code, err.AsMessage())
		return
	}
	utils.WriteJson(w, http.StatusOK, "Users populated")
}
