package app

import (
	"github.com/Dubjay18/green-lit/dto"
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
	for i, _ := range articles {
		articles[i].Content = utils.ParseRichText(articles[i].Content)
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
		article.Content = utils.ParseRichText(article.Content)
		utils.WriteJson(w, http.StatusOK, article)
	} else {
		utils.WriteJson(w, http.StatusBadRequest, "Invalid id")
	}

}

func (h ArticleHandler) GetArticlesByUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if i, err := strconv.Atoi(id); err != nil {
		articles, err := h.service.GetByUserID(i)
		if err != nil {
			utils.WriteJson(w, err.Code, err.AsMessage())
			return
		}
		for i, _ := range articles {
			articles[i].Content = utils.ParseRichText(articles[i].Content)
		}
		utils.WriteJson(w, http.StatusOK, articles)
	} else {
		utils.WriteJson(w, http.StatusBadRequest, "Invalid id")
	}
}

func (h ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var articleRequest dto.NewArticleRequest
	if err := utils.DecodeJson(r, &articleRequest); err != nil {
		utils.WriteJson(w, err.Code, err.AsMessage())
		return
	}
	resp, err := h.service.CreateArticle(articleRequest)
	if err != nil {
		utils.WriteJson(w, err.Code, err.AsMessage())
		return
	}
	utils.WriteJson(w, http.StatusCreated, resp)
}
