package handlers

import (
	"github.com/gin-gonic/gin"
	coinflipper "go-api-coin-flipper/domain/coin-flipper"
)

func HandleFlipCoin(context *gin.Context) {
	context.String(200, coinflipper.FlipCoin())
}
