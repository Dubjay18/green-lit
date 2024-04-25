package app

import (
	"fmt"
	"github.com/Dubjay18/green-lit/dto"
	"github.com/Dubjay18/green-lit/errs"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockAuthService struct{}

func (m MockAuthService) Login(request dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	//hp, _ := utils.HashPassword("password")
	//match := utils.CheckPasswordHash(request.Password, hp)
	fmt.Println("Login request: ", request.Email, request.Password)
	if request.Email == "test@gmail.com" && request.Password == "password" {
		return &dto.LoginResponse{
			AccessToken:  "valid-token",
			RefreshToken: "valid-refresh-token",
		}, nil
	}

	return nil, &errs.AppError{Code: http.StatusUnauthorized, Message: "Invalid credentials"}
}

func (m MockAuthService) Verify(urlParams map[string]string) *errs.AppError {
	if urlParams["token"] == "valid-token" {
		return nil
	}
	return &errs.AppError{Code: http.StatusUnauthorized, Message: "Invalid token"}
}

func TestAuthHandler_Login_BadRequestBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(`{"email":test@gmail.com,"password":`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := AuthHandler{service: MockAuthService{}}

	handler.Login(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "{\"message\":\"Invalid request payload\"}\n")
}

func TestAuthHandler_Login_EmptyEmailOrPassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(`{"email":"","password":"password"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := AuthHandler{service: MockAuthService{}}

	handler.Login(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid credentials")

	req, err = http.NewRequest("POST", "/auth/login", strings.NewReader(`{"email":"test@gmail.com","password":""}`))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.Login(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid credentials")
}
func TestAuthHandler_Verify(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/verify?token=valid-token", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := AuthHandler{service: MockAuthService{}}

	r := mux.NewRouter()
	r.HandleFunc("/auth/verify", handler.Verify)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "isAuthorized")

	req, err = http.NewRequest("GET", "/auth/verify?token=invalid-token", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid token")
}
