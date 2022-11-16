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

// Tweet represents a single Twitter post
type Tweet struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Language  string    `json:"lang"`
	Source    string    `json:"source"`
	Geo       TweetGeo  `json:"geo"`
}

// TweetGeo is a tweet location as a geometry (point in particular)
type TweetGeo struct {
	Coordinates Coordinate `json:"coordinates"`
}

// Coordinate is a coordinate object representing a point
type Coordinate struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
	PlaceID     string `json:"place_id"`
}
