package rest

type Client interface {
	Get(url string, opts ...GetOption) Request
}
