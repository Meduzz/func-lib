package main

import (
	"log"

	"github.com/Meduzz/func-lib/servicelib"
	"github.com/Meduzz/func-lib/servicelib/service"
	"github.com/Meduzz/func-lib/servicelib/service/annotation"
	"github.com/Meduzz/func-lib/servicelib/service/dto"
	"github.com/Meduzz/func-lib/servicelib/service/transports"
	"github.com/gin-gonic/gin"
)

func main() {
	normal := service.NewRole("normal", false)
	system := service.NewRole("system", true)

	trnsp := transports.Gin(
		"test.example.com",
		"",
	)

	trnsp.SetBefore(func() error {
		log.Println("This was the before hook")
		return nil
	})
	trnsp.GET("/hello/:world", func(ctx *gin.Context) {
		ctx.String(200, "Hello %s!", ctx.Param("world"))
	}, normal.Name)
	ep := trnsp.POST("/query", func(ctx *gin.Context) {
		bs, _ := ctx.GetRawData()

		log.Printf("%s", string(bs))

		ctx.JSON(200, gin.H{"name": "Query Result", "age": 42})
	}, system.Name)

	query := dto.NewField("query", dto.String)
	query.AddAnnotation(annotation.KeyValue("pattern", "\\w"))
	query.SetDescription("The query string")

	ep.SetDescription("This endpoint is used for callbacks, and this is markdown.")
	ep.SetExpects(dto.NewEntity("SearchQuery", dto.Fields(query)))
	ep.SetReturns(dto.NewEntity("SearchDocument", dto.Fields(
		dto.NewField("name", dto.String),
		dto.NewField("age", dto.Number),
	)))

	def := service.NewService(
		"test",
		"1.2.3",
		service.Envs("ASDF", "QWERTY"),
		service.APIs(trnsp),
		service.Roles(normal, system))

	def.SetDescription("A very long\npiece of text\ngoes here...")

	servicelib.Run(def)
}
