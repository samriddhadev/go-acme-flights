package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/api/validation/model"
)

func NewValidator() Validator {
	return Validator{}
}

type Validator struct {
}

func (validator *Validator) ValidateGetFlights(ctx *gin.Context) {
	query := model.GetFlightsQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}