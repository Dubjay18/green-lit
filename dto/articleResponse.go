package dto

type ArticleResponse struct {
	ID          int    `json:"article_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishedAt string `json:"published_at"`
	Author      string `json:"author"`
}
