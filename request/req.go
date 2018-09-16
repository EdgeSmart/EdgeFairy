package request

func (r *Request) Post(method string, pathinfo string, body interface{}) *Request {

	return r
}

func (r *Request) Do(method string, pathinfo string, body interface{}) *Response {

	return &Response{}
}
