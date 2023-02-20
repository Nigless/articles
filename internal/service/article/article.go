package article

import (
	"rest-api/internal/entity"
)

type ArticleRepo interface {
	Get(id int) (entity.Article, error)
}

type Article struct {
	repo ArticleRepo
}

func New(repo ArticleRepo) *Article {
	return &Article{repo}
}

func (self *Article) Get(id int) (entity.Article, error) {
	return self.repo.Get(id)
}
