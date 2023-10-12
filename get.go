package rest

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func Get(url string, opts ...Option) (*Result, error) {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	uri := generateURL(url, options.Queries)

	client := newClient(options.Insecure)
	defer client.CloseIdleConnections()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, nil)
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

func newClient(isSecure bool) *http.Client {
	if isSecure {
		return newSecureHTTPClient()
	}
	return newInsecureHTTPClient()
}

func newSecureHTTPClient() *http.Client {
	return &http.Client{}
}

func newInsecureHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec
			},
		},
	}
}

func generateURL(url string, queries map[string]string) string {
	if len(queries) == 0 {
		return url
	}
	strings := query(queries)
	wholeQueries := join(strings)
	return url + "?" + wholeQueries
}

func join(strings []string) string {
	if len(strings) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	head, tails := cutHead(strings)
	buffer.WriteString(head)
	for _, s := range tails {
		buffer.WriteString("&")
		buffer.WriteString(s)
	}
	return buffer.String()
}

func cutHead(strings []string) (string, []string) {
	return strings[0], strings[1:]
}

func query(queries map[string]string) []string {
	var ret []string
	for key, value := range queries {
		ret = append(ret, key+"="+url.QueryEscape(value))
	}
	return ret
}
