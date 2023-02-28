package http

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error"`
}

func ErrorResponse(ctx *gin.Context, err error) {
	log.Printf("%s", err)
	switch err {
	case sql.ErrNoRows:
		ctx.AbortWithStatus(404)

	default:
		ctx.AbortWithStatus(400)
	}

}
