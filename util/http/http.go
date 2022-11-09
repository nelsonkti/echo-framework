package http

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Http interface {
	Get(url string, params interface{}) ([]byte, error)
	Put(url string, params interface{}) ([]byte, error)
	Post(url string, params interface{}) ([]byte, error)
	Patch(url string, params interface{}) ([]byte, error)
	Delete(url string, params interface{}) ([]byte, error)
}

type http struct {
	req     *fasthttp.Request
	resp    *fasthttp.Response
	method  string
	options *options
}

func NewRequest(opts ...Option) Http {
	o := options{
		contentType: DefaultContentType,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &http{
		req:     fasthttp.AcquireRequest(),
		resp:    fasthttp.AcquireResponse(),
		options: &o,
	}
}

type Option func(*options)

type ContentType func() string

type options struct {
	contentType ContentType
}

func DefaultContentType() string {
	return "application/json"
}

func (h *http) Get(url string, params interface{}) ([]byte, error) {
	h.method = "GET"

	return h.request(url, params)
}

func (h *http) Put(url string, params interface{}) ([]byte, error) {
	h.method = "PUT"

	return h.request(url, params)
}

func (h *http) Post(url string, params interface{}) ([]byte, error) {
	h.method = "POST"

	return h.request(url, params)
}

func (h *http) Patch(url string, params interface{}) ([]byte, error) {
	h.method = "PATCH"

	return h.request(url, params)
}

func (h *http) Delete(url string, params interface{}) ([]byte, error) {
	h.method = "DELETE"

	return h.request(url, params)
}

func (h *http) request(url string, params interface{}) ([]byte, error) {
	h.req.SetRequestURI(url)
	h.req.Header.SetMethod(h.method)
	h.req.Header.SetContentType(h.options.contentType())

	buf, _ := json.Marshal(params)
	h.req.SetBody(buf)

	err := fasthttp.Do(h.req, h.resp)

	if err != nil {
		return nil, err
	}

	return h.resp.Body(), err
}

func WithContentType(contentType string) Option {
	return func(o *options) {
		o.contentType = func() string {
			return contentType
		}
	}
}
