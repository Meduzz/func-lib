package httplib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Meduzz/helper/http/client"
)

type (
	Test struct {
		Greeting string `json:"greeting,omitempty"`
		Name     string `json:"name,omitempty"`
		Age      int    `json:"age,omitempty"`
		Likes    string `json:"likes,omitempty"`
	}
)

func TestMain(m *testing.M) {

	Post("/test/:name", func(ctx *Context) {
		name := ctx.Param("name")
		age := ctx.QueryAsInt("age")
		likes := ctx.DefaultQuery("likes", "wine")
		test := &Test{}

		err := ctx.AsJSON(test)

		if err != nil {
			ctx.Text(500, err.Error(), "text/plain")
			return
		}

		test.Name = name
		test.Likes = likes
		test.Age = age

		ctx.SetHeader("X-Status", "Nailed it")

		ctx.JSON(200, test)
	})

	Get("/error", func(ctx *Context) {
		expect := ctx.Header("Expecting")
		body, err := ctx.Body()

		if err != nil {
			ctx.Text(500, err.Error(), "text/plain")
			return
		}

		if len(body) > 0 {
			ctx.Bytes(500, body, "text/plain")
			return
		}

		ctx.Text(200, expect, "text/plain")
	})

	Put("/", func(ctx *Context) {
		random := ctx.DefaultQuery("random", "0")

		body, _ := ctx.Body()

		if len(body) > 0 {
			test := &Test{}
			ctx.AsJSON(test)
		}

		ctx.Bytes(200, []byte(random), "text/plain")
	})

	Delete("/", func(ctx *Context) {
		ctx.Bytes(200, []byte("done"), "text/plain")
	})

	go Run()

	os.Exit(m.Run())
}

func TestPost(t *testing.T) {
	orig := &Test{
		Greeting: "Hello mr %s, you are %d and like %s!",
	}

	req, _ := client.POST("http://localhost:8080/test/Meduzz?age=9000", orig)

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
		return
	}

	if res.Code() != 200 {
		log.Printf("Response code was not 200 but %d\n", res.Code())
		t.FailNow()
		return
	}

	response := &Test{}
	res.Body(response)

	greeting := fmt.Sprintf(response.Greeting, response.Name, response.Age, response.Likes)

	if greeting != "Hello mr Meduzz, you are 9000 and like wine!" {
		log.Printf("Greeting was not correct, got %s\n", greeting)
		t.FailNow()
	}

	status := res.Header("X-Status")

	if status != "Nailed it" {
		log.Printf("Status was wrong, was %s\n", status)
		t.FailNow()
	}
}

func Test_Get(t *testing.T) {
	req, _ := client.GET("http://localhost:8080/error")
	req.Header("Expecting", "error")

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		fmt.Printf("Request threw error: %s\n", err.Error())
		t.FailNow()
	}

	if res.Code() != 200 {
		log.Printf("Code was not 200 but %d\n", res.Code())
		t.FailNow()
	}

	bs, _ := ioutil.ReadAll(res.Response().Body)

	if string(bs) != "error" {
		log.Printf("Body was not correct, was %s\n", string(bs))
		t.FailNow()
	}
}

func Test_Put(t *testing.T) {
	req, _ := client.PUT("http://localhost:8080/?random=1", &Test{Age: 10})

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		fmt.Printf("Request threw error: %s\n", err.Error())
		t.FailNow()
	}

	if res.Code() != 200 {
		log.Printf("Code was not 200 but %d\n", res.Code())
		t.FailNow()
	}

	bs, _ := ioutil.ReadAll(res.Response().Body)

	if string(bs) != "1" {
		log.Printf("Body was not correct, was %s\n", string(bs))
		t.FailNow()
	}
}

func Test_delete(t *testing.T) {
	req, _ := client.DELETE("http://localhost:8080/", nil)

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		fmt.Printf("Request threw error: %s\n", err.Error())
		t.FailNow()
	}

	if res.Code() != 200 {
		log.Printf("Code was not 200 but %d\n", res.Code())
		t.FailNow()
	}

	bs, _ := ioutil.ReadAll(res.Response().Body)

	if string(bs) != "done" {
		log.Printf("Body was not correct, was %s\n", string(bs))
		t.FailNow()
	}
}
