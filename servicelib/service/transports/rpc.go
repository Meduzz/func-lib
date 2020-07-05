package transports

import (
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/api"
)

type (
	RpcAPI struct {
		Name      string            `json:"name"`
		Type      string            `json:"type"`
		Endpoints []*RpcEndpointDTO `json:"endpoints"`
		Envs      []string          `json:"envs"`
	}

	RpcEndpointDTO struct {
		Topic   string `json:"topic"`
		Group   string `json:"group"`
		handler func(api.Context)
	}
)

func (r *RpcAPI) ApiType() string {
	return r.Type
}

func Rpc(name string) *RpcAPI {
	eps := make([]*RpcEndpointDTO, 0)
	envs := make([]string, 0)

	return &RpcAPI{name, "rpc", eps, envs}
}

func (r *RpcAPI) ApiName() string {
	return r.Name
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

func (r *RpcAPI) AddEnv(env string) {
	r.Envs = append(r.Envs, env)
}

func (r *RpcAPI) Envars() []string {
	return r.Envs
}
