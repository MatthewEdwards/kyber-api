package models

// Article is used to store a news article
type Article struct {
	Title string `json:"title"`
	Site  string `json:"site"`
	URL   string `json:"url"`
}
