// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"rest-api/internal/controller/http/v1/article"
	"rest-api/internal/service"

	"github.com/gin-gonic/gin"
)

func New(handler *gin.Engine, service *service.Service) {
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(404)
	})
	h := handler.Group("/v1")
	{
		article.New(h, &service.Article)
	}
}
