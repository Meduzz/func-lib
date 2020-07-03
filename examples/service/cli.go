package main

import (
	"log"

	"../../servicelib"
	"../../servicelib/service"
	"../../servicelib/service/transports"
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
