package transports

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
)

type (
	RpcAPI struct {
		ApiName   string            `json:"name"`
		ApiType   string            `json:"type"`
		Endpoints []*RpcEndpointDTO `json:"endpoints"`
	}

	RpcEndpointDTO struct {
		Topic   string `json:"topic"`
		Group   string `json:"group"`
		handler func(api.Context)
	}
)

func (r *RpcAPI) Type() string {
	return r.ApiType
}

func Rpc(name string) *RpcAPI {
	eps := make([]*RpcEndpointDTO, 0)

	return &RpcAPI{name, "rpc", eps}
}

func (r *RpcAPI) Name() string {
	return r.ApiName
}

func (r *RpcAPI) Handle(topic, group string, handler func(api.Context)) {
	r.Endpoints = append(r.Endpoints, &RpcEndpointDTO{topic, group, handler})
}

func (r *RpcAPI) Start() {
	conn, err := nuts.Connect()

	if err != nil {
		panic(err)
	}

	srv := rpc.NewRpc(conn)

	for _, ep := range r.Endpoints {
		srv.Handler(ep.Topic, ep.Group, ep.handler)
	}

	srv.Run()
}
