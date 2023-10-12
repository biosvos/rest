package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func Post(url string, opts ...Option) (*Result, error) {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	reader, err := newReader(options)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := newClient(options.Insecure)
	defer client.CloseIdleConnections()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, reader)
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

func newReader(options Options) (*bytes.Buffer, error) {
	if options.Data != nil {
		marshal, err := json.Marshal(options.Data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return bytes.NewBuffer(marshal), nil
	}
	if options.DataString != "" {
		return bytes.NewBufferString(options.DataString), nil
	}
	return bytes.NewBuffer(nil), nil
}
