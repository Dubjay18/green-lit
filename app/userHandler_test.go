package app

import (
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserService struct{}

func (m MockUserService) PopulateUsers() *errs.AppError {
	return nil
}

func (m MockUserService) GetAllUsers() ([]dto.UserResponse, *errs.AppError) {
	return []dto.UserResponse{{ID: 1, Email: "test@test.com", FullName: "Test User", Gender: "Male"}}, nil
}

func (m MockUserService) GetByID(id int) (*dto.UserResponse, *errs.AppError) {
	if id == 1 {
		return &dto.UserResponse{ID: 1, Email: "test@test.com", FullName: "Test User", Gender: "Male"}, nil
	}
	return nil, &errs.AppError{Code: http.StatusNotFound, Message: "User not found"}
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := UserHandler{service: MockUserService{}}

	handler.GetAllUsers(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `[{"user_id":1,"email":"test@test.com","full_name":"Test User","gender":"Male"}]
`, rr.Body.String())
}

func TestUserHandler_GetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := UserHandler{service: MockUserService{}}

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetUser)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"user_id":1,"email":"test@test.com","full_name":"Test User","gender":"Male"}
`, rr.Body.String())

	req, err = http.NewRequest("GET", "/users/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, `{"message":"User not found"}
`, rr.Body.String())
}
