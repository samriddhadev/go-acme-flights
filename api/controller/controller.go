package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/model"
	"github.com/samriddhadev/go-acme-flights/service"
)

func NewAcmeFlightController(flightService service.AcmeFlightService) AcmeFlightController {
	return AcmeFlightController{
		flightService: flightService,
	}
}

type AcmeFlightController struct {
	flightService service.AcmeFlightService
}

func (controller *AcmeFlightController) GetFlights(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		flights, err := controller.flightService.GetFlights(cfg)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, flights)
	}
}

func (controller *AcmeFlightController) CreateFlights(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var flight *model.Flight
		if err := ctx.BindJSON(&flight); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)	
		}
		err := controller.flightService.CreateFlight(cfg, flight)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusCreated, nil)
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