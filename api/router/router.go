package router

import (
	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/api/controller"
	"github.com/samriddhadev/go-acme-flights/config"
)

func NewRouter(cfg *config.Config, ctx *gin.Engine, controller controller.AcmeFlightController) Router {
	return Router{
		Cfg: cfg,
		Ctx: ctx,
		Controller: controller,
	}
}

type Router struct {
	Cfg *config.Config
	Ctx *gin.Engine
	Controller controller.AcmeFlightController
}

func (router *Router) Handle() {
	group := router.Ctx.Group(router.Cfg.BasePath)
	group.GET("/", router.Controller.GetFlights(router.Cfg))
	group.POST("/", router.Controller.CreateFlights(router.Cfg))
	group.GET("/:id", router.Controller.GetFlightById(router.Cfg))
	group.PUT("/:id", router.Controller.PutFlightById(router.Cfg))
	group.DELETE("/:id", router.Controller.DeleteFlightById(router.Cfg))
}