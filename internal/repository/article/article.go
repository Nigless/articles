package article

import (
	"database/sql"
	"rest-api/internal/entity"
)

type Article struct {
	*sql.DB
}

func New(db *sql.DB) *Article {
	return &Article{db}
}

func (self *Article) Get(id int) (entity.Article, error) {
	var article entity.Article
	return article, self.QueryRow("SELECT * FROM article WHERE id=$1", id).Scan(&article)
}
