//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/samriddhadev/go-acme-flights/repository"
	"github.com/samriddhadev/go-acme-flights/service"
	"github.com/samriddhadev/go-acme-flights/api/validation"
	"github.com/samriddhadev/go-acme-flights/api/controller"
	"github.com/samriddhadev/go-acme-flights/api/router"
	"github.com/samriddhadev/go-acme-flights/config"
)

func InitializeRouter(cfg *config.Config, ctx *gin.Engine) router.Router {
	wire.Build(
		router.NewRouter, 
		controller.NewAcmeFlightController,
		validation.NewValidator,
		service.NewAcmeFlightService,
		repository.NewAcmeFlightRepository,
	)
	return router.Router{}
}
