package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	apperror "github.com/samriddhadev/go-acme-flights/api/error"
)

func ErrorHandlerMiddleware(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		for _, ginErr := range ctx.Errors {
			switch err := ginErr.Err.(type) {
			case *apperror.ValidationError:
				log.Printf("validation: error - %s", err.Error())
				ctx.JSON(-1, err)
				return
			case *apperror.NonFatalApiError:
				log.Printf("non-fatal: error - %s", err.Error())
			case *apperror.FatalApiError:
				log.Printf("fatal: error - %s", err.Error())
				ctx.JSON(-1, err)
				return
			case error:
				log.Printf("runtime: error - %s", err.Error())
				ctx.JSON(http.StatusInternalServerError, apperror.FatalApiError{Message: ginErr.Err.Error()})
				return
			}
		}
	}
}