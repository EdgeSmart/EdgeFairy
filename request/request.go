package request

import (
	"net/http"
)

type Request struct {
	http.Request
}

func Post(url string, data interface{}) *Request {
	return &Request{}
}

func Get(url string, data interface{}) *Request {
	return &Request{}
}

func Put(url string, data interface{}) *Request {
	return &Request{}
}

func Head(url string, data interface{}) *Request {
	return &Request{}
}

func Delete(url string, data interface{}) *Request {
	return &Request{}
}

func Options(url string, data interface{}) *Request {
	return &Request{}
}
