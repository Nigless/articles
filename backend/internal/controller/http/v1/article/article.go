package article

import (
	"rest-api/internal/controller/http"
	"rest-api/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	Get(id int64) (entity.Article, error)
	Create(article entity.Article) (int64, error)
}

type Article struct {
	articles ArticleService
}

func New(handler *gin.RouterGroup, service ArticleService) *Article {
	article := &Article{service}
	handler.GET("article/:id", article.Get)
	handler.POST("article", article.Create)
	return article
}

func (self *Article) Get(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	article, err := self.articles.Get(id)
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	ctx.AbortWithStatusJSON(200, article)
}

func (self *Article) Create(ctx *gin.Context) {
	var article entity.Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	id, err := self.articles.Create(article)
	if err != nil {
		http.ErrorResponse(ctx, err)
		return
	}

	ctx.AbortWithStatusJSON(200, id)
}
