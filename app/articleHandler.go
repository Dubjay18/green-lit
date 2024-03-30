package app

import (
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	service service.ArticleService
}

func (h ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		utils.WriteJson(w, err.Code, err.AsMessage())
		return
	}
	utils.WriteJson(w, http.StatusOK, articles)

}

func (h ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if i, err := strconv.Atoi(id); err != nil {
		article, err := h.service.GetByID(i)
		if err != nil {
			utils.WriteJson(w, err.Code, err.AsMessage())
			return
		}
		utils.WriteJson(w, http.StatusOK, article)
	} else {
		utils.WriteJson(w, http.StatusBadRequest, "Invalid id")
	}

}
