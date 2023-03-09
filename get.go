package rest

type GetOptions struct {
	Queries map[string]string
	Headers map[string]string
}

type GetOption func(options *GetOptions)

func WithQueries(queries map[string]string) GetOption {
	return func(options *GetOptions) {
		options.Queries = queries
	}
}

func WithHeaders(headers map[string]string) GetOption {
	return func(options *GetOptions) {
		options.Headers = headers
	}
}

func ApplyGetOptions(opts []GetOption) GetOptions {
	var options GetOptions
	for _, opt := range opts {
		opt(&options)
	}
	return options
}
