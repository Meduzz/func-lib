package service

import (
	"github.com/Meduzz/func-lib/servicelib/service/transports"
)

type (
	Role struct {
		Name    string `json:"name"`
		Private bool   `json:"private"`
	}

	ServiceDefinitionDTO struct {
		Name        string                      `json:"name"`
		Version     string                      `json:"version"`
		Description string                      `json:"description"`
		Envs        []string                    `json:"envs"`
		APIs        []*transports.ApiDefinition `json:"api"`
		Roles       []*Role                     `json:"roles"`
	}
)

func NewService(name, version string, envs []string, apis []*transports.ApiDefinition, roles []*Role) *ServiceDefinitionDTO {
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

func APIs(apis ...*transports.ApiDefinition) []*transports.ApiDefinition {
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
