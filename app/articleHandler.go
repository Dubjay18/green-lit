package app

import (
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"net/http"
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
