package arvancloud

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func buildURI(path string, options interface{}) string {
	v, _ := query.Values(options)
	return (&url.URL{Path: path, RawQuery: v.Encode()}).String()
}
