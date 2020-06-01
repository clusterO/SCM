package service

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

/* Calling other services */

// proxymw implements AuthService, forwarding Authenticate requests to the
// provided endpoint, and serving all other (i.e. Count) requests via the
// next AuthService.
type proxymw struct {
	next      AuthService     // Serve most requests via this service...
	authenticate endpoint.Endpoint // ...except Authenticate, which gets served by this endpoint
}

/* Client-side endpoints */
func (mw proxymw) Authenticate(username string, password string) (string, error) {
	response, err := mw.authenticate(authenticateRequest{S: username})
	if err != nil {
		return "", err
	}
	resp := response.(authenticateResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func proxyingMiddleware(proxyURL string) ServiceMiddleware {
	return func(next AuthService) AuthService {
		return proxymw{next, makeUppercaseProxy(proxyURL)}
	}
}

func makeAuthenticateProxy(proxyURL string) endpoint.Endpoint {
	return httptransport.NewClient(
		"GET",
		mustParseURL(proxyURL),
		encodeUppercaseRequest,
		decodeUppercaseResponse,
	).Endpoint()
}

/* Service discovery and load balancing */