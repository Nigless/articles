package article

import (
	"rest-api/internal/entity"
)

type ArticleRepo interface {
	Get(id int64) (entity.Article, error)
	Create(article entity.Article) (int64, error)
}

type Article struct {
	repo ArticleRepo
}

func New(repo ArticleRepo) *Article {
	return &Article{repo}
}

func (self *Article) Get(id int64) (entity.Article, error) {
	return self.repo.Get(id)
}

func (self *Article) Create(article entity.Article) (int64, error) {
	return self.repo.Create(article)
}
