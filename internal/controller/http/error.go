package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error"`
}

func ErrorResponse(ctx *gin.Context, err error) {
	switch err {
	case sql.ErrNoRows:
		ctx.AbortWithStatus(404)

	default:
		ctx.AbortWithStatus(400)
	}

}
