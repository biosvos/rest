package http

import "github.com/biosvos/rest"

var _ rest.Client = &Client{}

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (n *Client) Get(url string, opts ...rest.GetOption) rest.Request {
	options := rest.ApplyGetOptions(opts)
	url = generateUrl(url, options.Queries)
	return NewRequest(url)
}
