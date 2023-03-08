package http

import (
	"bytes"
	"net/url"
)

func generateUrl(url string, queries map[string]string) string {
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
