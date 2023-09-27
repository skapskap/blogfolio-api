package data

import (
	"database/sql"
	"errors"
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
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT id, created_at, updated_at, published_at, title, description, status
	FROM blog
	WHERE id = $1`

	var post Post

	err := p.DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.PublishedAt,
		&post.Title,
		&post.Description,
		&post.Status,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &post, nil
}

func (p PostModel) Update(post *Post) error {
	query := `
		UPDATE blog
		SET title = $1, updated_at = NOW(), description = $2, status = $3
		WHERE id = $4`

	args := []interface{}{
		post.Title,
		post.Description,
		post.Status,
		post.ID,
	}

	_, err := p.DB.Exec(query, args...)
	return err
}

func (p PostModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE from blog
		WHERE id = $1`

	result, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (p PostModel) GetAll() ([]*Post, error) {
	query := `
		SELECT id, created_at, updated_at, published_at, title, description, status
		FROM blog`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.PublishedAt,
			&post.Title,
			&post.Description,
			&post.Status,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
