package repository

import (
	"rest-api/internal/repository/article"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	*article.Article
}

func New(db *sqlx.DB) *Repository {
	return &Repository{article.New(db)}
}
