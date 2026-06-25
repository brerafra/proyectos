package domain

type User struct {
	UserId   int64  `json:"id"`
	Name     string `json:"name"`
	Card     int64  `json:"card"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}
