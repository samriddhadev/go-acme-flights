package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdHandlerMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("x-request-id", uuid.New().String())
	ctx.Next()
}