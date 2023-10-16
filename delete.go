package rest

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func Delete(url string, opts ...Option) (*Result, error) {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	uri := generateURL(url, options.Queries)

	client := newClient(options.Insecure)
	defer client.CloseIdleConnections()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, uri, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header = options.Headers

	rsp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		err := rsp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	contents, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Result{
		Header:   rsp.Header,
		Code:     rsp.StatusCode,
		Contents: contents,
	}, nil
}
