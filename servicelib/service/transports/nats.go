package transports

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/nats-io/nats.go"
)

type (
	NatsAPI struct {
		Name      string             `json:"name"`
		Type      string             `json:"type"`
		Envs      []string           `json:"envs"`
		Endpoints []*NatsEndpointDTO `json:"endpoints"`
	}

	NatsEndpointDTO struct {
		Topic   string `json:"topic"`
		Group   string `json:"group,omitempty"`
		handler nats.MsgHandler
	}
)

func (n *NatsAPI) ApiType() string {
	return n.Type
}

func (n *NatsAPI) ApiName() string {
	return n.Name
}

func (n *NatsAPI) Envars() []string {
	return n.Envs
}

func Nats(name string) *NatsAPI {
	envs := make([]string, 0)
	eps := make([]*NatsEndpointDTO, 0)

	n := &NatsAPI{
		Name:      name,
		Type:      "nats",
		Endpoints: eps,
		Envs:      envs,
	}

	n.AddEnv("NATS_URL")

	return n
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
	conn, err := nuts.Connect()

	if err != nil {
		return err
	}

	for _, ep := range n.Endpoints {
		if ep.Group != "" {
			conn.QueueSubscribe(ep.Topic, ep.Group, ep.handler)
		} else {
			conn.Subscribe(ep.Topic, ep.handler)
		}
	}

	return nil
}

func (n *NatsAPI) AddEnv(env string) {
	n.Envs = append(n.Envs, env)
}
