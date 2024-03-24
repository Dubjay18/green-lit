package domain

// User represents a user in the system
type User struct {
	ID       int    `json:"id"db:"id"csv:"id"`
	FullName string `json:"full_name"db:"full_name"csv:"full_name"`
	Email    string `json:"email"db:"email"csv:"email"`
	Password string `json:"password"db:"password"csv:"password"`
	//Articles []userArticle `json:"articles"db:"articles"csv:"articles"`
	Gender string `json:"gender"db:"gender"csv:"gender"`
}

type userArticle struct {
	ID int `json:"id"db:"id"csv:"id"`
}
