package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/api/controller"
	"github.com/samriddhadev/go-acme-flights/api/middleware"
	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/logger"
)

func NewRouter(
	cfg *config.Config,
	ctx *gin.Engine,
	logger *logger.AcmeLogger,
	errorMiddleware *middleware.ErrorMiddleware,
	requestIdMiddleware *middleware.RequestIdMiddleware,
	controller controller.AcmeFlightController,
) Router {
	return Router{
		Cfg:                 cfg,
		Ctx:                 ctx,
		logger:              logger,
		errorMiddleware:     errorMiddleware,
		requestIdMiddleware: requestIdMiddleware,
		Controller:          controller,
	}
}

type Router struct {
	Cfg                 *config.Config
	Ctx                 *gin.Engine
	logger              *logger.AcmeLogger
	errorMiddleware     *middleware.ErrorMiddleware
	requestIdMiddleware *middleware.RequestIdMiddleware
	Controller          controller.AcmeFlightController
}

func (router *Router) Handle() {
	group := router.Ctx.Group(router.Cfg.BasePath)
	group.Use(router.errorMiddleware.Apply, router.requestIdMiddleware.Apply)
	group.GET("/", router.Controller.GetFlights(router.Cfg))
	group.POST("/", router.Controller.CreateFlight(router.Cfg))
	group.GET("/:id", router.Controller.GetFlightById(router.Cfg))
	group.PUT("/:id", router.Controller.PutFlightById(router.Cfg))
	group.DELETE("/:id", router.Controller.DeleteFlightById(router.Cfg))

	health := router.Ctx.Group(router.Cfg.HealthCheck)
	health.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
}
