package gateway

import (
	"net/http"
)

type Gateway struct {
	RESTGateway
}

func NewGateway() *Gateway {
	c := &http.Client{
		// Timeout: time.Second * 30,
	}
	return &Gateway{
		NewRESTGateway(c),
	}
}
