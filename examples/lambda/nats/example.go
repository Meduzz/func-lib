package main

import (
	"fmt"

	"github.com/Meduzz/func-lib/lambda"
	"github.com/nats-io/go-nats"
)

func main() {
	lambda.Nats(func(msg *nats.Msg) {
		text := string(msg.Data)
		fmt.Printf("Got message: %s\n", text)
	})
}
