package controller

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"sync"

	"exam.com/webchecker/internal/core"
	"exam.com/webchecker/internal/entities"
	"github.com/gin-gonic/gin"
)

type WebCheckerController interface {
	CheckAvalibleWebsites(*gin.Context)
}

type webcheckerController struct {
	webCheckerCore core.WebCheckerCore
}

func NewWebCheckerController(webCheckerCore core.WebCheckerCore) WebCheckerController {
	return &webcheckerController{webCheckerCore: webCheckerCore}
}

type request struct {
	entities.Websites `json:"websites"`
}

type response struct {
	WebsiteStatusList []entities.WebsiteStatus `json:"website_status_list"`
}

// CheckAvalibleWebsites - Read a list of websites in an uploaded file and follow the processes to determine whether or not they are reachable.
func (c *webcheckerController) CheckAvalibleWebsites(ctx *gin.Context) {
	res := response{}

	file, _, err := ctx.Request.FormFile("csv")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	chInputWebsite := make(chan string)
	chOutputWebsiteStatus := make(chan entities.WebsiteStatus)

	go func() {
		csvReader := csv.NewReader(file)
		for {
			record, err := csvReader.Read()
			if err == io.EOF || len(record) == 0 {
				break
			}

			chInputWebsite <- record[0]
		}
		close(chInputWebsite)

	}()

	c.process(chInputWebsite, chOutputWebsiteStatus, &wg)

	go func() {
		wg.Wait()
		close(chOutputWebsiteStatus)
	}()

	for output := range chOutputWebsiteStatus {
		go func(output entities.WebsiteStatus) {
			defer wg.Done()
			res.WebsiteStatusList = append(res.WebsiteStatusList, output)
		}(output)
	}
	ctx.JSON(http.StatusOK, res)

}

func (c *webcheckerController) process(chInputWebsite <-chan string, chOutputWebsiteStatus chan<- entities.WebsiteStatus, wg *sync.WaitGroup) error {
	for website := range chInputWebsite {
		wg.Add(1)
		go func(website string) {
			websiteStatus, err := c.webCheckerCore.CheckAvalibleWebsite(website)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			chOutputWebsiteStatus <- websiteStatus
		}(website)
	}
	return nil
}
