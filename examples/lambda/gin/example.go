package main

import (
	"github.com/Meduzz/func-lib/lambda"
	"github.com/gin-gonic/gin"
)

func main() {
	lambda.Gin(func(srv *gin.Engine) {
		srv.POST("/http", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"Hello": "world!"})
		})
	})
}
