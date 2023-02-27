package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samriddhadev/go-acme-flights/core/logger"
)

func NewRequestIdMiddleware(logger *logger.AcmeLogger) *RequestIdMiddleware {
	return &RequestIdMiddleware {
		logger: logger,
	}
}

type RequestIdMiddleware struct {
	logger *logger.AcmeLogger
}

func (middleware *RequestIdMiddleware) Apply(ctx *gin.Context) {
	ctx.Writer.Header().Set("x-request-id", uuid.New().String())
	ctx.Next()
}