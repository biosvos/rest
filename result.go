package rest

import "fmt"

type Result struct {
	Header   map[string][]string
	Code     int
	Contents []byte
}

func (r *Result) String() string {
	return fmt.Sprintf("%v %v %v", r.Header, r.Code, string(r.Contents))
}
