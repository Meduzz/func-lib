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
		Name        string         `json:"name"`
		Version     string         `json:"version"`
		Description string         `json:"description"`
		Envs        []string       `json:"envs"`
		APIs        map[string]API `json:"api"`
		Roles       []*Role        `json:"roles"`
	}
)

func NewService(name, version string, envs []string, apis []API, roles []*Role) *ServiceDefinitionDTO {
	apiMap := make(map[string]API)

	for _, api := range apis {
		apiMap[api.ApiName()] = api
	}

	return &ServiceDefinitionDTO{
		Name:    name,
		Version: version,
		Envs:    envs,
		APIs:    apiMap,
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

func (s *ServiceDefinitionDTO) SetDescription(desc string) {
	s.Description = desc
}
