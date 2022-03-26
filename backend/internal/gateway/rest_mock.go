package gateway

import (
	"math/rand"
	"net/http"
	"time"
)

type mockRESTGateway struct {
	err error
}

// NewMockRESTGateway - This mock gateway is used for unit testing to make sure it does not make a real request.
func NewMockRESTGateway(err error) RESTGateway {
	return &mockRESTGateway{err: err}
}

// MakeRequest - This mock has a random amount of time to sleep for testing concurrency.
// If it does not give an error, it always gives the status 200.
func (g *mockRESTGateway) MakeGETRequest(url string) (*http.Response, error) {
	r := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(r))

	if g.err != nil {
		return nil, g.err
	}

	res := &http.Response{StatusCode: 200}

	return res, nil
}
