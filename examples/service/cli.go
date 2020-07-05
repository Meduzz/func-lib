package main

import (
	"log"

	"github.com/Meduzz/func-lib/servicelib"
	"github.com/Meduzz/func-lib/servicelib/service"
	"github.com/Meduzz/func-lib/servicelib/service/transports"
	"github.com/gin-gonic/gin"
)

func main() {
	normal := service.NewRole("normal", false)
	system := service.NewRole("system", true)

	trnsp := transports.Gin(
		"test.example.com",
		"/test",
		"/on",
	)

	trnsp.GET("/hello/:world", func(ctx *gin.Context) {
		ctx.String(200, "Hello %s!", ctx.Param("world"))
	}, normal.Name)
	ep := trnsp.POST("/on", func(ctx *gin.Context) {
		bs, _ := ctx.GetRawData()

		log.Printf("%s", string(bs))

		ctx.Status(200)
	}, system.Name)
	trnsp.SetBefore(func() error {
		log.Println("This was the before hook")
		return nil
	})

	ep.SetDescription("This endpoint is used for callbacks, and this is markdown.")
	ep.SetRequestEncoding("application/json")
	ep.SetResponseEncoding("text/plain")
	ep.SetExpects(gin.H{"type": "fake", "some": "value"})
	ep.SetReturns(gin.H{"type": "fake", "some": "value"})

	def := service.NewService(
		"test",
		"1.2.3",
		service.Envs("ASDF", "QWERTY"),
		service.APIs(trnsp),
		service.Roles(normal, system))

	def.SetDescription("A very long\nmarkdown text\ngoes here...")

	servicelib.Run(def)
}
