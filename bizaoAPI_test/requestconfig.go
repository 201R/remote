package remote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

//Options is to set up the request.
type Options struct {
	//URL to visit
	URL string

	//data(map) to query on a domain
	Query map[string][]string

	//Request header
	Header map[string]string

	//Request body
	Body interface{}
}

// Parse parses raw url
func (config *Options) parse() (string, io.Reader) {

	u, err := url.Parse(config.URL)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	for k, v := range config.Query {
		for _, t := range v {
			q.Set(k, t)
		}
	}

	u.RawQuery = q.Encode()

	b, err := json.Marshal(config.Body)
	if err != nil || config.Body == nil {
		return u.String(), nil
	}

	return u.String(), bytes.NewReader(b)
}
