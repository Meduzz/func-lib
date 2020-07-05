package transports

import (
	"log"

	"github.com/Meduzz/helper/utilz"
	"github.com/gin-gonic/gin"
)

type (
	GinAPI struct {
		srv       *gin.Engine
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
		Method string   `json:"method"`
		URL    string   `json:"url"`
		Roles  []string `json:"roles"`
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
	gin.SetMode(gin.ReleaseMode)

	srv := gin.Default()
	endpoints := make([]*EndpointDTO, 0)
	envs := []string{"PORT"}
	empty := func() error { return nil }

	return &GinAPI{srv, empty, "gin", "http", domain, context, cbURL, endpoints, envs}
}

func (g *GinAPI) GET(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.GET(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"GET", url, roles})
}

func (g *GinAPI) POST(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.POST(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"POST", url, roles})
}

func (g *GinAPI) PUT(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.PUT(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"PUT", url, roles})
}

func (g *GinAPI) DELETE(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.DELETE(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"DELETE", url, roles})
}

func (g *GinAPI) HEAD(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.HEAD(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"HEAD", url, roles})
}

func (g *GinAPI) OPTIONS(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.OPTIONS(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"OPTIONS", url, roles})
}

func (g *GinAPI) PATCH(url string, handler gin.HandlerFunc, roles ...string) {
	g.srv.PATCH(url, handler)
	g.Endpoints = append(g.Endpoints, &EndpointDTO{"PATCH", url, roles})
}

func (g *GinAPI) Start() error {
	err := g.before()

	if err != nil {
		return err
	}

	port := utilz.Env("PORT", ":8080")

	log.Printf("Starting gin on port: %s\n", port)
	return g.srv.Run(port)
}

func (g *GinAPI) SetBefore(hook func() error) {
	g.before = hook
}
