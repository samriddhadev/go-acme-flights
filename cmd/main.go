package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/di"
)

const HOST string = "localhost"

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf("error starting server - %s", err)
	}
	router := di.InitializeRouter(cfg, engine)
	router.Handle()
	log.Printf("server started at %s:%d", HOST, cfg.Port)
	engine.Run(fmt.Sprintf("%s:%s", HOST, strconv.Itoa(cfg.Port)))
}
