package transports

type (
	API interface {
		Start() error
	}

	ApiDefinition struct {
		Type       string   `json:"type"`
		Name       string   `json:"name"`
		Envs       []string `json:"envs"`
		Definition API      `json:"definition"`
	}
)

func Start(api *ApiDefinition) error {
	return api.Definition.Start()
}

func (r *ApiDefinition) AddEnv(env string) {
	r.Envs = append(r.Envs, env)
}
