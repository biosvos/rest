package rest

type Options struct {
	Queries    map[string]string
	Headers    map[string][]string
	Insecure   bool
	Data       any
	DataString string
}

type Option func(options *Options)

func WithQueries(queries map[string]string) Option {
	return func(options *Options) {
		options.Queries = queries
	}
}

func WithHeaders(headers map[string][]string) Option {
	return func(options *Options) {
		options.Headers = headers
	}
}

func WithInsecure() Option {
	return func(options *Options) {
		options.Insecure = true
	}
}

func WithData(a any) Option {
	return func(options *Options) {
		options.Data = a
	}
}

func WithDataString(data string) Option {
	return func(options *Options) {
		options.DataString = data
	}
}
