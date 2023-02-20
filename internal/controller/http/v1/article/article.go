package article

import (
	"rest-api/internal/controller/http"
	"rest-api/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	Get(id int) (entity.Article, error)
}

type Article struct {
	articles ArticleService
}

func New(handler *gin.RouterGroup, service ArticleService) *Article {
	article := &Article{service}
	handler.GET("article/:id", article.Get)
	return article
}

func (self *Article) Get(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	article, err := self.articles.Get(id)
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	ctx.SecureJSON(200, article)
}
