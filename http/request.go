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
	url string
}

func NewRequest(url string) *Request {
	return &Request{url: url}
}

func (r *Request) Execute() ([]byte, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return all, nil
}
