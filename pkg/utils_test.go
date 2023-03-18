package arvancloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Pagination struct {
	Page    int `json:"page,omitempty" url:"page,omitempty"`
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
}

type testExample struct {
	A string `url:"a,omitempty"`
	B string `url:"b,omitempty"`
	Pagination
}

func Test_buildURI(t *testing.T) {
	tests := map[string]struct {
		path   string
		params interface{}
		want   string
	}{
		"multi level":                  {path: "/accounts/test.ir", params: testExample{}, want: "/accounts/test.ir"},
		"multi level - params":         {path: "/domains/test.ir", params: testExample{A: "b"}, want: "/domains/test.ir?a=b"},
		"multi level - multi params":   {path: "/domains/test.ir", params: testExample{A: "b", B: "d"}, want: "/domains/test.ir?a=b&b=d"},
		"multi level - nested fields":  {path: "/domains/test.ir", params: testExample{A: "b", B: "d", Pagination: Pagination{PerPage: 10}}, want: "/domains/test.ir?a=b&b=d&per_page=10"},
		"single level":                 {path: "/test", params: testExample{}, want: "/test"},
		"single level - params":        {path: "/test", params: testExample{B: "d"}, want: "/test?b=d"},
		"single level - multi params":  {path: "/test", params: testExample{A: "b", B: "d"}, want: "/test?a=b&b=d"},
		"single level - nested fields": {path: "/test", params: testExample{A: "b", B: "d", Pagination: Pagination{PerPage: 10}}, want: "/test?a=b&b=d&per_page=10"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := buildURI(tc.path, tc.params)
			assert.Equal(t, tc.want, got)
		})
	}
}
