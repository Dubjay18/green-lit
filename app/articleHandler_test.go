package app

import (
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockArticleService struct{}

func (s *MockArticleService) GetByID(id int) (*dto.ArticleResponse, *errs.AppError) {
	return &dto.ArticleResponse{
		ID:          "1",
		Title:       "Title",
		Content:     "Content",
		PublishedAt: "2021-01-01T00:00:00Z",
		Author:      1,
	}, nil
}

func (s *MockArticleService) GetByUserID(userID int) ([]dto.ArticleResponse, *errs.AppError) {
	return []dto.ArticleResponse{
		{
			ID:          "1",
			Title:       "Title",
			Content:     "Content",
			PublishedAt: "2021-01-01T00:00:00Z",
			Author:      1,
		},
	}, nil
}

func (s *MockArticleService) CreateArticle(req dto.NewArticleRequest) (*dto.ArticleResponse, *errs.AppError) {
	return &dto.ArticleResponse{
		ID:          "1",
		Title:       req.Title,
		Content:     req.Content,
		PublishedAt: "2021-01-01T00:00:00Z",
		Author:      1,
	}, nil
}

func (s *MockArticleService) GetAllArticles() ([]dto.ArticleResponse, *errs.AppError) {
	return []dto.ArticleResponse{
		{
			ID:          "1",
			Title:       "Title",
			Content:     "Content",
			PublishedAt: "2021-01-01T00:00:00Z",
			Author:      1,
		},
	}, nil
}

func TestGetAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &ArticleHandler{service: &MockArticleService{}}

	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.GetAllArticles).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"article_id":"1","title":"Title","content":"Content","published_at":"2021-01-01T00:00:00Z","author":1}]`
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetArticle(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &ArticleHandler{service: &MockArticleService{}}

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", handler.GetArticle).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"article_id":"1","title":"Title","content":"Content","published_at":"2021-01-01T00:00:00Z","author":1}`
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateArticle(t *testing.T) {
	article := `{"title":"Title","content":"Content","author":1}`
	req, err := http.NewRequest("POST", "/articles", strings.NewReader(article))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &ArticleHandler{service: &MockArticleService{}}

	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := `{"article_id":"1","title":"Title","content":"Content","published_at":"2021-01-01T00:00:00Z","author":1}`
	actual := strings.TrimSuffix(rr.Body.String(), "\n")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
