package repository

import (
	"database/sql"
	"rest-api/internal/repository/article"
)

type Repository struct {
	*article.Article
}

func New(db *sql.DB) *Repository {
	return &Repository{article.New(db)}
}
