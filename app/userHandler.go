package app

import (
	"fmt"
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		utils.WriteJson(w, err.Code, err.AsMessage())
		return
	}
	utils.WriteJson(w, http.StatusOK, users)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if i, err := strconv.Atoi(id); err == nil {
		user, err := h.service.GetByID(i)
		if err != nil {
			utils.WriteJson(w, err.Code, err.AsMessage())
			return
		}
		utils.WriteJson(w, http.StatusOK, user)
	} else {
		utils.WriteJson(w, http.StatusBadRequest, fmt.Sprintf("Invalid id: %s", id))
	}

}
