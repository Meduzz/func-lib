package natslib

import (
	"encoding/json"
	"time"

	"github.com/Meduzz/helper/nuts"
	nats "github.com/nats-io/go-nats"
)

type (
	NatsLib struct {
		conn *nats.Conn
	}
)

func NewNatsLib() (*NatsLib, error) {
	conn, err := nuts.Connect()

	if err != nil {
		return nil, err
	}

	lib := &NatsLib{conn}

	return lib, nil
}

func (l *NatsLib) Subscribe(topic, queue string, handler nats.MsgHandler) (*nats.Subscription, error) {
	if queue == "" {
		return l.conn.Subscribe(topic, handler)
	} else {
		return l.conn.QueueSubscribe(topic, queue, handler)
	}
}

func (l *NatsLib) RequestBytes(topic string, data []byte, deadline int) ([]byte, error) {
	msg, err := l.conn.Request(topic, data, time.Duration(deadline)*time.Second)

	if err != nil {
		return nil, err
	}

	return msg.Data, nil
}

func (l *NatsLib) RequestJson(topic string, request interface{}, reply interface{}, deadline int) error {
	jsonData, err := json.Marshal(request)

	if err != nil {
		return err
	}

	msg, err := l.conn.Request(topic, jsonData, time.Duration(deadline)*time.Second)

	if err != nil {
		return err
	}

	return json.Unmarshal(msg.Data, reply)
}

func (l *NatsLib) PublishBytes(topic string, data []byte) error {
	return l.conn.Publish(topic, data)
}

func (l *NatsLib) PublishJson(topic string, data interface{}) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return l.conn.Publish(topic, jsonData)
}
