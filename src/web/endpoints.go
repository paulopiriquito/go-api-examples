package web

import (
	"github.com/gin-gonic/gin"
	"go-api-coin-flipper/web/handlers"
)

func RegisterEndpoints(router *gin.Engine) gin.IRoutes {
	router.GET("/healthcheck", handlers.Healthcheck)

	addCoinFlipper(router)

	return router
}

func addCoinFlipper(router *gin.Engine) gin.IRoutes {
	return router.GET("/coin-flip", handlers.HandleFlipCoin)
}
