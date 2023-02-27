package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperror "github.com/samriddhadev/go-acme-flights/api/error"
	"github.com/samriddhadev/go-acme-flights/core/logger"
)

func NewErrorMiddleware(logger *logger.AcmeLogger) *ErrorMiddleware {
	return &ErrorMiddleware{
		logger: logger,
	}
}

type ErrorMiddleware struct {
	logger *logger.AcmeLogger
}

func (middleware *ErrorMiddleware) Apply(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		for _, ginErr := range ctx.Errors {
			switch err := ginErr.Err.(type) {
			case *apperror.ValidationError:
				middleware.logger.Errorf("validation: error - %s", err.Error())
				ctx.JSON(-1, err)
				return
			case *apperror.NonFatalApiError:
				middleware.logger.Errorf("non-fatal: error - %s", err.Error())
			case *apperror.FatalApiError:
				middleware.logger.Errorf("fatal: error - %s", err.Error())
				ctx.JSON(-1, err)
				return
			case error:
				middleware.logger.Errorf("runtime: error - %s", err.Error())
				ctx.JSON(http.StatusInternalServerError, apperror.FatalApiError{Message: ginErr.Err.Error()})
				return
			}
		}
	}
}