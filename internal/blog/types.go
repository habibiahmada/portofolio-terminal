package blog

// Post is a published blog article returned by the public API.
type Post struct {
	ID                 string    `json:"id"`
	Slug               string    `json:"slug"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	BodyMD             string    `json:"body_md"`
	Category           string    `json:"category"`
	Tags               []string  `json:"tags"`
	ReadingTimeMinutes int       `json:"reading_time_minutes"`
	PublishedAt        string   `json:"published_at"`
}

// listResponse matches the paginated envelope from the website public API.
type listResponse struct {
	Success bool   `json:"success"`
	Data    []Post `json:"data"`
}

// postsEnvelope is a fallback shape when the API returns a bare array.
type postsEnvelope []Post
