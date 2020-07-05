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
		"/callback",
	)

	trnsp.GET("/hello/:world", func(ctx *gin.Context) {
		ctx.String(200, "Hello %s!", ctx.Param("world"))
	}, normal.Name)
	trnsp.POST("/on", func(ctx *gin.Context) {
		bs, _ := ctx.GetRawData()

		log.Printf("%s", string(bs))

		ctx.Status(200)
	}, system.Name)
	trnsp.SetBefore(func() error {
		log.Println("This was the before hook")
		return nil
	})

	def := service.NewService(
		"test",
		"1.2.3",
		service.Envs("ASDF", "QWERTY"),
		service.APIs(trnsp),
		service.Roles(normal, system))

	servicelib.Run(def)
}
