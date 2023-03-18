package arvancloud

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func buildURI(path string, options interface{}) string {
	v, _ := query.Values(options)
	return (&url.URL{Path: path, RawQuery: v.Encode()}).String()
}

// PrettyPrint will print the contents of the object
func PrettyPrint(data interface{}) {
	p, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s \n", p)
}
