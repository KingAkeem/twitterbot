package store

import "time"

// User represents a single Twitter user
type User struct {
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	ProfileImageURL string    `json:"profile_image_url"`
	URL             string    `json:"url"`
	Verified        bool      `json:"verified"`
}
