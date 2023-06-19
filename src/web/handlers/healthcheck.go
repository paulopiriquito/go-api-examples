package handlers

import "github.com/gin-gonic/gin"

func Healthcheck(context *gin.Context) {
	context.Data(200, "text/html; charset=utf-8", []byte("healthy"))
}
