package httplib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	Context struct {
		request  *http.Request
		response http.ResponseWriter
		params   httprouter.Params
		body     []byte
		status   int
	}
)

func (c *Context) AsJSON(to interface{}) error {
	bs, err := c.Body()

	if err != nil {
		return err
	}

	return json.Unmarshal(bs, to)
}

func (c *Context) AsString() (string, error) {
	bs, err := c.Body()

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func (c *Context) Body() ([]byte, error) {
	if len(c.body) == 0 {
		bs, err := ioutil.ReadAll(c.request.Body)

		if err != nil {
			return nil, err
		}

		c.body = bs

		return bs, nil
	}

	return c.body, nil
}

func (c *Context) JSON(status int, body interface{}) error {
	bs, err := json.Marshal(body)

	if err != nil {
		return err
	}

	c.response.Header().Add("Content-Type", "application/json")
	c.response.WriteHeader(status)
	c.status = status
	_, err = c.response.Write(bs)

	return err
}

func (c *Context) Text(status int, text string, contentType string) error {
	c.response.Header().Add("Content-Type", contentType)
	c.response.WriteHeader(status)
	c.status = status
	_, err := c.response.Write([]byte(text))

	return err
}

func (c *Context) Bytes(status int, body []byte, contentType string) error {
	c.response.Header().Add("Content-Type", contentType)
	c.response.WriteHeader(status)
	c.status = status
	_, err := c.response.Write(body)

	return err
}

func (c *Context) Header(name string) string {
	return c.request.Header.Get(name)
}

func (c *Context) SetHeader(name, value string) {
	c.response.Header().Add(name, value)
}

func (c *Context) Param(name string) string {
	return c.params.ByName(name)
}

func (c *Context) DefaultQuery(key, value string) string {
	q := c.request.URL.Query().Get(key)

	if q != "" {
		return q
	}

	return value
}

func (c *Context) QueryAsInt(key string) int {
	q := c.request.URL.Query().Get(key)

	if q != "" {
		i, err := strconv.Atoi(q)

		if err != nil {
			return 0
		}

		return i
	}

	return 0
}

func (c *Context) Logf(format string, params ...string) {
	l := fmt.Sprintf(format, params)
	log.Printf("%s\n", l)
}
