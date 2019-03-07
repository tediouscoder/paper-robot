package http

import (
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
)

// Head is the http Head request.
func Head(url string) (*http.Response, error) {
	return client.Head(url)
}
