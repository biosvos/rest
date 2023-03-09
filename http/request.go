package http

import (
	"github.com/biosvos/rest"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
)

var _ rest.Request = &Request{}

type Request struct {
	method  string
	url     string
	headers map[string]string
}

func NewRequest(method, url string, headers map[string]string) *Request {
	return &Request{
		method:  method,
		url:     url,
		headers: headers,
	}
}

func (r *Request) Execute() ([]byte, error) {
	req, err := createHttpRequest(r.method, r.url, r.headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret, err := doRequest(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

func createHttpRequest(method string, url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	setHeaders(req, headers)
	return req, nil
}

func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
