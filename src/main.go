package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-api-coin-flipper/web"
	"go-api-coin-flipper/web/middleware"
)

func main() {
	router := gin.New()
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(gin.Recovery())
	web.RegisterEndpoints(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal().Msg("Service stopped listening and is shutting down")
	}
}
