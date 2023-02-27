package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/di"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf("error starting server - %s", err)
	}
	router := di.InitializeRouter(cfg, engine)
	router.Handle()
	log.Printf("server started at localhost:%d", cfg.Port)
	engine.Run(fmt.Sprintf(":%s", strconv.Itoa(cfg.Port)))
}
