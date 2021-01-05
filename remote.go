package remote

import (
	"net/http"
	"time"
)

//Content type
const (
	contentJSON = "application/json"
	contentFORM = "application/x-www-form-urlencoded"
)

var nbr int64

type (
	//Remote Client
	Remote struct {
		client *http.Client
	}

	//Config is to set up the client option
	Config struct {
		//add cookie jar on avery request
		Jar http.CookieJar

		//Request time out
		Timeout time.Duration

		//check if there are a script to redirect the request on another URL
		CheckRedirect func(req *http.Request, via []*http.Request) error
	}
)

//NewRemote create a client to do a request to a specifique URL
func NewRemote(config Config) *Remote {
	client := &http.DefaultClient

	return &Remote{
		client: *client,
	}
}

//GET Makes GET requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) GET(config Options) (*http.Response, error) {
	return remote.do(http.MethodGet, contentFORM, config)
}

//HEAD Makes HEAD requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) HEAD(config Options) (*http.Response, error) {
	return remote.do(http.MethodHead, contentFORM, config)
}

//POST Makes POST requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) POST(config Options) (*http.Response, error) {
	return remote.do(http.MethodPost, contentJSON, config)
}

//PUT Makes PUT requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) PUT(config Options) (*http.Response, error) {
	return remote.do(http.MethodPut, contentJSON, config)
}

//PATCH Makes PATCH requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) PATCH(config Options) (*http.Response, error) {
	return remote.do(http.MethodPatch, contentJSON, config)
}

//DELETE Makes DELETE requests on a specific url
//with the configurations defined in RequestConfig
func (remote *Remote) DELETE(config Options) (*http.Response, error) {
	return remote.do(http.MethodDelete, contentJSON, config)
}

func (remote *Remote) do(method string, content string, config Options) (*http.Response, error) {
	url, body := config.parse()

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	for k, v := range config.Header {
		request.Header.Add(k, v)
	}

	request.Header.Set("Content-Type", content)

	//time.Sleep(2 * time.Minute)
	res, err := remote.client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
