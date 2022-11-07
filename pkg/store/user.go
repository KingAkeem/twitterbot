package store

// User represents a single Twitter user
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	ID       string `json:"id"`
}
