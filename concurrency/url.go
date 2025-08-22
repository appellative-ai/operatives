package concurrency

import (
	"net/http"
	"net/url"
	"strings"
)

func newURL(req *http.Request, host string) error {
	if req == nil {
		return nil
	}
	scheme := "https"
	if strings.Contains(host, "localhost:") {
		scheme = "http"
	}
	u, err := url.Parse(scheme + "://" + host + req.URL.String())
	if err != nil {
		return err
	}
	req.URL = u
	req.Host = host
	return nil
}
