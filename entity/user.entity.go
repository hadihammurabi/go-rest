package entity

// User entity
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
