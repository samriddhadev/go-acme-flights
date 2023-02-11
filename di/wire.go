//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/samriddhadev/go-acme-flights/api/controller"
	"github.com/samriddhadev/go-acme-flights/api/router"
	"github.com/samriddhadev/go-acme-flights/config"
)

func InitializeRouter(cfg *config.Config, ctx *gin.Engine) router.Router {
	wire.Build(
		router.NewRouter, 
		controller.NewAcmeFlightController,
	)
	return router.Router{}
}
