package transports

import (
	"../annotation"
	"../dto"
	"github.com/Meduzz/helper/utilz"
	"github.com/gin-gonic/gin"
)

type (
	GinAPI struct {
		before    func() error
		Name      string            `json:"name"`
		Type      string            `json:"type"`
		Domain    string            `json:"domain"`
		Context   string            `json:"context"`
		Endpoints []*GinEndpointDTO `json:"endpoints"`
		Envs      []string          `json:"envs"`
	}

	GinEndpointDTO struct {
		Method      string   `json:"method"`
		URL         string   `json:"url"`
		Roles       []string `json:"roles"`
		handler     gin.HandlerFunc
		Description string                  `json:"description,omitempty"`
		Expects     *dto.EntityDTO          `json:"expects,omitempty"`
		Returns     *dto.EntityDTO          `json:"returns,omitempty"`
		Annotations []annotation.Annotation `json:"annotations,omitempty"`
	}
)

func (g *GinAPI) ApiType() string {
	return g.Type
}

func (g *GinAPI) ApiName() string {
	return g.Name
}

func (g *GinAPI) Envars() []string {
	return g.Envs
}

func Gin(domain, context string) *GinAPI {
	endpoints := make([]*GinEndpointDTO, 0)
	envs := []string{"PORT"}
	empty := func() error { return nil }

	return &GinAPI{empty, "gin", "http", domain, context, endpoints, envs}
}

func (g *GinAPI) GET(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "GET",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) POST(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "POST",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) PUT(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "PUT",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) DELETE(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "DELETE",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) HEAD(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "HEAD",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) OPTIONS(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "OPTIONS",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) PATCH(url string, handler gin.HandlerFunc, roles ...string) *GinEndpointDTO {
	ep := &GinEndpointDTO{
		Method:  "PATCH",
		URL:     url,
		Roles:   roles,
		handler: handler,
	}
	g.Endpoints = append(g.Endpoints, ep)

	return ep
}

func (g *GinAPI) Start() error {
	err := g.before()

	if err != nil {
		return err
	}

	srv := gin.Default()

	for _, ep := range g.Endpoints {
		switch ep.Method {
		case "GET":
			srv.GET(ep.URL, ep.handler)
		case "POST":
			srv.POST(ep.URL, ep.handler)
		case "PUT":
			srv.PUT(ep.URL, ep.handler)
		case "DELETE":
			srv.DELETE(ep.URL, ep.handler)
		case "HEAD":
			srv.HEAD(ep.URL, ep.handler)
		case "OPTIONS":
			srv.OPTIONS(ep.URL, ep.handler)
		case "PATCH":
			srv.PATCH(ep.URL, ep.handler)
		}
	}

	port := utilz.Env("PORT", ":8080")

	return srv.Run(port)
}

func (g *GinAPI) SetBefore(hook func() error) {
	g.before = hook
}

func (e *GinEndpointDTO) SetDescription(desc string) {
	e.Description = desc
}

func (e *GinEndpointDTO) SetExpects(entity *dto.EntityDTO) {
	e.Expects = entity
}

func (e *GinEndpointDTO) SetReturns(entity *dto.EntityDTO) {
	e.Returns = entity
}
