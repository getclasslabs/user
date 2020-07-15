package pkg

import "net/http"

//HttpClientInterface interface to mock http client
type HttpClientInterface interface{
	Do(req *http.Request) (*http.Response, error)
}