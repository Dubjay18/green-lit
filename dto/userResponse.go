package dto

type UserResponse struct {
	ID       int    `json:"user_id"db:"id"`
	Email    string `json:"email"db:"email"`
	FullName string `json:"full_name"db:"full_name"`
	Gender   string `json:"gender"db:"gender"`
}
