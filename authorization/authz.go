package authorization

import (
	"github.com/appellative-ai/common/core"
	"net/http"
)

const (
	HandlerName = "common:resiliency:handler/authorization/http"
	AuthzName   = "Authorization"
)

func Authorization(next core.Exchange) core.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(AuthzName)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		return next(r)
	}
}
