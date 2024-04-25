package dto

type LoginRequest struct {
	Email    string `json:"email"db:"email"`
	Password string `json:"password"`
}
