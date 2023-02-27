package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/api/validation"
	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/logger"
	"github.com/samriddhadev/go-acme-flights/model"
	"github.com/samriddhadev/go-acme-flights/service"
)

func NewAcmeFlightController(logger *logger.AcmeLogger, validator validation.Validator, flightService service.AcmeFlightService) AcmeFlightController {
	return AcmeFlightController{
		logger: logger,
		validator: validator,
		flightService: flightService,
	}
}

type AcmeFlightController struct {
	logger *logger.AcmeLogger
	validator validation.Validator
	flightService service.AcmeFlightService
}

func (controller *AcmeFlightController) GetFlights(cfg *config.Config) gin.HandlerFunc {
	return controller.validator.Validate(validation.SCHEMA_GET_FLIGHTS, cfg, func(ctx *gin.Context) {
		var flightFilter model.FlightFilter
		if err := ctx.ShouldBind(&flightFilter); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		flights, err := controller.flightService.GetFlights(cfg, &flightFilter)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, flights)
	})
}

func (controller *AcmeFlightController) CreateFlight(cfg *config.Config) gin.HandlerFunc {
	return controller.validator.Validate(validation.SCHEMA_CREATE_FLIGHT, cfg, func(ctx *gin.Context) {
		var flight *model.Flight
		if err := ctx.BindJSON(&flight); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		err := controller.flightService.CreateFlight(cfg, flight)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusCreated, nil)
	})
}

func (controller *AcmeFlightController) GetFlightById(cfg *config.Config) gin.HandlerFunc {
	return controller.validator.Validate(validation.SCHEMA_GET_FLIGHT, cfg, func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		flight, err := controller.flightService.GetFlight(cfg, id)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, flight)
	})
}

func (controller *AcmeFlightController) PutFlightById(cfg *config.Config) gin.HandlerFunc {
	return controller.validator.Validate(validation.SCHEMA_UPDATE_FLIGHT, cfg, func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var input *model.Flight
		if err := ctx.BindJSON(&input); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		flight, err := controller.flightService.UpdateFlight(cfg, id, input)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, flight)
	})
}

func (controller *AcmeFlightController) DeleteFlightById(cfg *config.Config) gin.HandlerFunc {
	return controller.validator.Validate(validation.SCHEMA_DELETE_FLIGHT, cfg, func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		err = controller.flightService.DeleteFlight(cfg, id)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	})
}