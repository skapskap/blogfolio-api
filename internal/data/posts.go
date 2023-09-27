package data

import (
	"database/sql"
	"time"
)

type Post struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

type PostModel struct {
	DB *sql.DB
}

func (p PostModel) Insert(post *Post) error {
	query := `
	INSERT INTO blog (title, description, status)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at, published_at
`

	args := []interface{}{post.Title, post.Description, post.Status}

	return p.DB.QueryRow(query, args...).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.PublishedAt)
}

func (p PostModel) Get(id int64) (*Post, error) {
	return nil, nil
}

func (p PostModel) Update(post *Post) error {
	return nil
}

func (p PostModel) Delete(id int64) error {
	return nil
}
