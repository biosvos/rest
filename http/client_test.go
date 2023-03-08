package http

import (
	"github.com/stretchr/testify/require"
	"rest"
	"testing"
)

func TestHttpGet(t *testing.T) {
	body, err := NewClient().Get("https://map.kakao.com/route/pubtrans.json", rest.WithQueries(map[string]string{
		"inputCoordSystem":  "WCONGNAMUL",
		"outputCoordSystem": "WCONGNAMUL",
		"service":           "map.daum.net",
		"sX":                "508505",
		"sY":                "1103177",
		"sName":             "시민의숲.양재꽃시장(양곡도매시장+방면)",
		"sid":               "BS69978",
		"eX":                "496570",
		"eY":                "1069229",
		"eName":             "의왕톨게이트(경기도인재개발원.경기연구원.경기도평생교육진흥원.경기도여성가족재단+방면)",
		"eid":               "BS81753",
	})).Execute()
	require.NoError(t, err)
	t.Log(body)
}
