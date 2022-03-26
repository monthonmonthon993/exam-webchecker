package gateway

import (
	"net/http"
)

type restGateway struct {
	client *http.Client
}

type RESTGateway interface {
	MakeGETRequest(url string) (*http.Response, error)
}

func NewRESTGateway(client *http.Client) RESTGateway {
	return &restGateway{client: client}
}

// MakeGETRequest - Makes a request to the given url and returns the status code and any errors encountered.
func (g *restGateway) MakeGETRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
