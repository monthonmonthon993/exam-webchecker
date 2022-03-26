package core

import (
	"fmt"

	"exam.com/webchecker/internal/entities"
	"exam.com/webchecker/internal/gateway"
	"exam.com/webchecker/pkg/null"
)

type WebCheckerCore interface {
	CheckAvalibleWebsite(string) (entities.WebsiteStatus, error)
}

type webCheckerCore struct {
	websiteGateway gateway.RESTGateway
}

func NewWebCheckerCore(websiteGateway gateway.RESTGateway) WebCheckerCore {
	return &webCheckerCore{websiteGateway: websiteGateway}
}

// CheckAvalibleWebsites - makes a request by website url and checks the reachable status of the website.
// returns a status code and an error.
func (w *webCheckerCore) CheckAvalibleWebsite(website string) (entities.WebsiteStatus, error) {
	websiteStatus := entities.WebsiteStatus{Name: website}
	if res, err := w.websiteGateway.MakeGETRequest(website); err != nil {
		fmt.Printf("Down: cannot reach the %s, error: %s\n", website, err.Error())
		websiteStatus.ErrMsg = null.NewString(err.Error())
		websiteStatus.OK = false
	} else {
		fmt.Printf("Up: http status %d that is returned from the %s\n", res.StatusCode, website)
		websiteStatus.OK = true
	}
	return websiteStatus, nil
}
