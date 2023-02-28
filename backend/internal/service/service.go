package service

import (
	"rest-api/internal/repository"
	"rest-api/internal/service/article"
)

type Service struct {
	article.Article
}

func New(repo *repository.Repository) *Service {
	return &Service{
		*article.New(repo.Article),
	}
}
