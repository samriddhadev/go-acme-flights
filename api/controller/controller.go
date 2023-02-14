package controller

import (
	"log"
	"net/http"
	"strconv"

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
			return
		}
		ctx.JSON(http.StatusOK, flights)
	}
}

func (controller *AcmeFlightController) CreateFlight(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var flight *model.Flight
		if err := ctx.BindJSON(&flight); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return	
		}
		err := controller.flightService.CreateFlight(cfg, flight)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusCreated, nil)
	}
}

func (controller *AcmeFlightController) GetFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		flight, err := controller.flightService.GetFlight(cfg, id)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, flight)
	}
}

func (controller *AcmeFlightController) PutFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		var input *model.Flight
		if err := ctx.BindJSON(&input); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return	
		}
		flight, err := controller.flightService.UpdateFlight(cfg, id, input)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, flight)
	}
}

func (controller *AcmeFlightController) DeleteFlightById(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = controller.flightService.DeleteFlight(cfg, id)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}