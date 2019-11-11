package lambda

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
)

// Gin - starts a gin webserver based on what's setup in the setup lambda.
func Gin(setup func(*gin.Engine)) error {
	port := flag.String("bind", "0.0.0.0:8080", "Set what ip&port this server will bind to.")
	flag.Parse()

	srv := gin.Default()
	setup(srv)

	/*
		TODO
		Verify signatures + autosetup of keys
	*/

	return srv.Run(*port)
}

// Nats - Create a nats listener server from the provided lambda.
func Nats(handler func(*nats.Msg)) error {
	host := flag.String("host", "127.0.0.1:4222", "nats host to connect to")
	topic := flag.String("topic", "", "nats topic to bind to.")
	queue := flag.String("queue", "", "queue group to share load with.")

	flag.Parse()

	if *topic == "" {
		return fmt.Errorf("No topic was provided")
	}

	conn, err := nats.Connect(fmt.Sprintf("nats://%s", *host))

	if err != nil {
		return err
	}

	if *queue != "" {
		conn.QueueSubscribe(*topic, *queue, handler)
	} else {
		conn.Subscribe(*topic, handler)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	conn.Close()

	return nil
}

// RPC - start a rpc server from the provided lambda.
func RPC(handler func(api.Context)) error {
	host := flag.String("host", "nats://127.0.0.1:4222", "nats host to connect to")
	topic := flag.String("topic", "", "nats topic to bind to.")
	queue := flag.String("queue", "", "queue group to share load with.")

	flag.Parse()

	if *topic == "" {
		return fmt.Errorf("No topic was provided")
	}

	conn, err := nats.Connect(*host)

	if err != nil {
		return err
	}

	srv := rpc.NewRpc(conn)

	err = srv.Handler(*topic, *queue, handler)

	if err != nil {
		return err
	}

	srv.Run()

	/*
		TODO
	*/

	return nil
}

/*
	TODO
	Add handler for net or tcp/udp. (handler gets a connection)
	Add handler for web (plain golang http server)
	Add handler for grpc... because reasons...
*/
