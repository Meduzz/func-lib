package main

import (
	"github.com/Meduzz/func-lib/lambda"
	"github.com/Meduzz/rpc/api"
)

func main() {
	lambda.RPC(func(ctx api.Context) {
		body, _ := ctx.Body()
		ctx.Reply(body)
	})
}
