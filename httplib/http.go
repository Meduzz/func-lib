package httplib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type (
	HandlerFunc func(*Context)
)

var router = httprouter.New()

func adapter(handler HandlerFunc) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		method := method(req.Method)
		path := req.RequestURI
		start := time.Now()

		ctx := &Context{
			request:  req,
			response: res,
			params:   params,
		}

		handler(ctx)
		duration := time.Now().Sub(start)

		log.Printf("%s %s - %d (%s)", method, path, ctx.status, duration.String())
	}
}

func Post(path string, handler HandlerFunc) {
	router.POST(path, adapter(handler))
}

func Get(path string, handler HandlerFunc) {
	router.GET(path, adapter(handler))
}

func Put(path string, handler HandlerFunc) {
	router.PUT(path, adapter(handler))
}

func Delete(path string, handler HandlerFunc) {
	router.DELETE(path, adapter(handler))
}

func Head(path string, handler HandlerFunc) {
	router.HEAD(path, adapter(handler))
}

func Options(path string, handler HandlerFunc) {
	router.OPTIONS(path, adapter(handler))
}

func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func method(m string) string {
	if len(m) < len("options") {
		return method(fmt.Sprintf("%s ", m))
	}

	return m
}
