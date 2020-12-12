package transports

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/nats-io/nats.go"
)

type (
	NatsAPI struct {
		Endpoints []*NatsEndpointDTO `json:"endpoints"`
		conn      *nats.Conn
	}

	NatsEndpointDTO struct {
		Topic   string `json:"topic"`
		Group   string `json:"group,omitempty"`
		handler nats.MsgHandler
	}
)

func Nats(name string) (*ApiDefinition, *NatsAPI) {
	envs := []string{"NATS_URL"}
	eps := make([]*NatsEndpointDTO, 0)

	n := &NatsAPI{
		Endpoints: eps,
	}

	return &ApiDefinition{
		Name:       name,
		Type:       "nats",
		Envs:       envs,
		Definition: n,
	}, n
}

func (n *NatsAPI) SetConn(conn *nats.Conn) {
	n.conn = conn
}

func (n *NatsAPI) Handle(topic, group string, handler nats.MsgHandler) {
	ep := &NatsEndpointDTO{
		Topic:   topic,
		Group:   group,
		handler: handler,
	}

	n.Endpoints = append(n.Endpoints, ep)
}

func (n *NatsAPI) Start() error {
	if n.conn == nil {
		conn, err := nuts.Connect()

		if err != nil {
			return err
		}

		n.conn = conn
	}

	for _, ep := range n.Endpoints {
		if ep.Group != "" {
			n.conn.QueueSubscribe(ep.Topic, ep.Group, ep.handler)
		} else {
			n.conn.Subscribe(ep.Topic, ep.handler)
		}
	}

	return nil
}
