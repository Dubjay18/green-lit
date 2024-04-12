package dto

type ArticleResponse struct {
	ID          string `json:"article_id" db:"article_id"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
	PublishedAt string `json:"published_at" db:"published_at"`
	Author      int    `json:"author" db:"author_id"`
}
