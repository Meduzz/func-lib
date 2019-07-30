package natslib

import (
	"log"
	"os"
	"testing"

	nats "github.com/nats-io/go-nats"
)

type Test struct {
	Msg string `json:"msg,omitempty"`
}

var natsLib *NatsLib

func TestMain(m *testing.M) {
	lib, err := NewNatsLib()

	if err != nil {
		log.Fatalln("Could not connect to nats")
		os.Exit(1)
	}

	natsLib = lib

	_, err = lib.Subscribe("natslib.echo", "", func(msg *nats.Msg) {
		if msg.Reply != "" {
			lib.PublishBytes(msg.Reply, msg.Data)
		} else {
			log.Println(string(msg.Data))
		}
	})

	if err != nil {
		log.Fatalf("Could not create subscription %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func Test_Request(t *testing.T) {
	reply, err := natsLib.RequestBytes("natslib.echo", []byte("Hello world"), 3)

	if err != nil {
		log.Printf("Request threw error: %s\n", err.Error())
		t.FailNow()
	}

	if string(reply) != "Hello world" {
		log.Printf("Respone was wrong, was %s\n", string(reply))
		t.FailNow()
	}
}

func Test_Request_Json(t *testing.T) {
	msg := &Test{"Hello world"}
	reply := &Test{}

	err := natsLib.RequestJson("natslib.echo", msg, reply, 3)

	if err != nil {
		log.Printf("Request threw error: %s\n", err.Error())
		t.FailNow()
	}

	if reply.Msg != "Hello world" {
		log.Printf("Respone was wrong, was %s\n", reply.Msg)
		t.FailNow()
	}
}

func Test_Publish(t *testing.T) {
	err := natsLib.PublishBytes("natslib.echo", []byte("Hello world"))

	if err != nil {
		log.Printf("Public threw error: %s\n", err.Error())
		t.FailNow()
	}
}

func Test_Publish_Json(t *testing.T) {
	msg := &Test{"Hello world"}

	err := natsLib.PublishJson("natslib.echo", msg)

	if err != nil {
		log.Printf("Public threw error: %s\n", err.Error())
		t.FailNow()
	}
}
