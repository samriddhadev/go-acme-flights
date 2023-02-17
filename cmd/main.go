package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/api/middleware"
	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/di"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf("error starting server - %s", err)
	}
	engine.Use(middleware.RequestIdHandlerMiddleware, middleware.ErrorHandlerMiddleware)
	router := di.InitializeRouter(cfg, engine)
	router.Handle()
	log.Printf("server started at localhost:%d", cfg.Port)
	engine.Run(fmt.Sprintf(":%s", strconv.Itoa(cfg.Port)))
}
