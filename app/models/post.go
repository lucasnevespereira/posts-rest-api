package models

type Post struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type JsonPost struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
