package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	host   string
	port   string
	client http.Client
	header map[string]string
	cookie map[string]string
}

func NewRequest(host string, port string) Request {
	return Request{
		client: http.Client{},
		host:   host,
		port:   port,
	}
}

func (r *Request) Request(method string, pathinfo string, body interface{}) {
	req, err := http.NewRequest("POST", r.host+pathinfo, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}
	r.client.
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := r.client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func (r *Request) POST() {

}

func (r *Request) Do() {

}
