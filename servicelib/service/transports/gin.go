package transports

import (
	"github.com/Meduzz/helper/utilz"
	"github.com/gin-gonic/gin"
)

type (
	GinAPI struct {
		before    func() error
		Name      string         `json:"name"`
		Type      string         `json:"type"`
		Domain    string         `json:"domain"`
		Context   string         `json:"context"`
		CbURL     string         `json:"callback"`
		Endpoints []*EndpointDTO `json:"endpoints"`
		Envs      []string       `json:"envs"`
	}

	EndpointDTO struct {
		Method  string   `json:"method"`
		URL     string   `json:"url"`
		Roles   []string `json:"roles"`
		handler gin.HandlerFunc
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

func Gin(domain, context, cbURL string) *GinAPI {
	endpoints := make([]*EndpointDTO, 0)
	envs := []string{"PORT"}
	empty := func() error { return nil }

	return &GinAPI{empty, "gin", "http", domain, context, cbURL, endpoints, envs}
}

func (g *GinAPI) GET(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"GET", url, roles, handler})
}

func (g *GinAPI) POST(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"POST", url, roles, handler})
}

func (g *GinAPI) PUT(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"PUT", url, roles, handler})
}

func (g *GinAPI) DELETE(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"DELETE", url, roles, handler})
}

func (g *GinAPI) HEAD(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"HEAD", url, roles, handler})
}

func (g *GinAPI) OPTIONS(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"OPTIONS", url, roles, handler})
}

func (g *GinAPI) PATCH(url string, handler gin.HandlerFunc, roles ...string) {
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"PATCH", url, roles, handler})
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
