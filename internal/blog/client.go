package blog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const defaultAPIBase = "https://www.habibiahmada.dev"

// Client fetches published blog posts from the public website API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a client using HABIBIAHMADA_API_BASE or the default origin.
func NewClient() *Client {
	base := strings.TrimRight(os.Getenv("HABIBIAHMADA_API_BASE"), "/")
	if base == "" {
		base = defaultAPIBase
	}
	return &Client{
		BaseURL: base,
		HTTPClient: &http.Client{
			Timeout: 12 * time.Second,
		},
	}
}

// FetchPublished returns all published posts, newest first.
func (c *Client) FetchPublished() ([]Post, error) {
	url := c.BaseURL + "/api/public/blog?page_size=50"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("blog fetch: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blog fetch: HTTP %d", resp.StatusCode)
	}

	posts, err := decodePosts(body)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func decodePosts(body []byte) ([]Post, error) {
	var wrapped listResponse
	if err := json.Unmarshal(body, &wrapped); err == nil && wrapped.Success && len(wrapped.Data) > 0 {
		return wrapped.Data, nil
	}
	if err := json.Unmarshal(body, &wrapped); err == nil && len(wrapped.Data) > 0 {
		return wrapped.Data, nil
	}

	var direct postsEnvelope
	if err := json.Unmarshal(body, &direct); err == nil && len(direct) > 0 {
		return direct, nil
	}

	var obj struct {
		Items []Post `json:"items"`
	}
	if err := json.Unmarshal(body, &obj); err == nil && len(obj.Items) > 0 {
		return obj.Items, nil
	}

	return nil, fmt.Errorf("blog fetch: unrecognized response")
}

// FormatDate renders a post date for the TUI list/detail header.
func FormatDate(raw string) string {
	if raw == "" {
		return ""
	}
	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05Z07:00", "2006-01-02"} {
		if t, err := time.Parse(layout, raw); err == nil {
			return t.Format("2 Jan 2006")
		}
	}
	return raw
}

// CategoryLabel returns a human-readable category name.
func CategoryLabel(category string) string {
	labels := map[string]string{
		"programming":       "Programming",
		"education":         "Education",
		"web":               "Web",
		"career":            "Career",
		"opinion":           "Opinion",
		"news-commentary":   "News",
	}
	if l, ok := labels[category]; ok {
		return l
	}
	return category
}
