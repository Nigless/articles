// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"rest-api/internal/controller/http/v1/article"
	"rest-api/internal/service"
	"rest-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func New(handler *gin.Engine, service *service.Service, logger logger.Logger) {
	h := handler.Group("/v1")
	{
		article.New(h, &service.Article)
	}
}
