package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Rest struct {
	Handler *gin.Engine
	log     zerolog.Logger
}

func New(log zerolog.Logger) *Rest {
	handler := gin.New()
	self := &Rest{
		handler,
		log,
	}

	handler.Use(self.logger)
	return self
}

func (self *Rest) logger(ctx *gin.Context) {
	duration := time.Since(time.Now())

	ctx.Next()
	log := self.log.
		Info().
		Str("method", ctx.Request.Method).
		Str("path", ctx.Request.URL.Path).
		Int("status", ctx.Writer.Status()).
		Dur("duration", duration)

	log.Msg("http")
}
