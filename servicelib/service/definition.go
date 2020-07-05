package service

type (
	API interface {
		ApiType() string
		ApiName() string
		Start() error
		Envars() []string
	}

	Role struct {
		Name    string `json:"name"`
		Private bool   `json:"private"`
	}

	ServiceDefinitionDTO struct {
		Name    string   `json:"name"`
		Version string   `json:"version"`
		Envs    []string `json:"envs"`
		APIs    []API    `json:"api"`
		Roles   []*Role  `json:"roles"`
	}
)

func NewService(name, version string, envs []string, apis []API, roles []*Role) *ServiceDefinitionDTO {
	return &ServiceDefinitionDTO{
		Name:    name,
		Version: version,
		Envs:    envs,
		APIs:    apis,
		Roles:   roles,
	}
}

func Envs(envs ...string) []string {
	return envs
}

func APIs(apis ...API) []API {
	return apis
}

func Roles(roles ...*Role) []*Role {
	return roles
}

func NewRole(name string, private bool) *Role {
	return &Role{name, private}
}
