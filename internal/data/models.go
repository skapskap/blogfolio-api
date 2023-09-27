package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("registro não encontrado")

type Models struct {
	Posts PostModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Posts: PostModel{DB: db},
	}
}
