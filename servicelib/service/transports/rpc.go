package transports

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
)

type (
	RpcAPI struct {
		Endpoints []*RpcEndpointDTO `json:"endpoints"`
	}

	RpcEndpointDTO struct {
		Topic   string `json:"topic"`
		Group   string `json:"group"`
		handler func(api.Context)
	}
)

func Rpc(name string) (*ApiDefinition, *RpcAPI) {
	eps := make([]*RpcEndpointDTO, 0)
	envs := make([]string, 0)

	api := &RpcAPI{eps}

	return &ApiDefinition{
		Name:       name,
		Type:       "rpc",
		Envs:       envs,
		Definition: api,
	}, api
}

func (r *RpcAPI) Handle(topic, group string, handler func(api.Context)) {
	r.Endpoints = append(r.Endpoints, &RpcEndpointDTO{topic, group, handler})
}

func (r *RpcAPI) Start() error {
	conn, err := nuts.Connect()

	if err != nil {
		return err
	}

	srv := rpc.NewRpc(conn)

	for _, ep := range r.Endpoints {
		srv.Handler(ep.Topic, ep.Group, ep.handler)
	}

	srv.Run()

	return nil
}
