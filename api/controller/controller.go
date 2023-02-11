package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/config"
)

func NewAcmeFlightController() AcmeFlightController {
	return AcmeFlightController{}
}

type AcmeFlightController struct {

}

func (controller *AcmeFlightController) GetFlights(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (controller *AcmeFlightController) CreateFlights(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (controller *AcmeFlightController) GetFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (controller *AcmeFlightController) PutFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (controller *AcmeFlightController) DeleteFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}