package gateway

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_restGateway_MakeRequest(t *testing.T) {
	wantStatusCode := 200

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.AbortWithError(wantStatusCode, errors.New("cannot reach to this site"))
	})
	server := httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
	defer server.Close()

	g := &restGateway{client: server.Client()}
	got, err := g.MakeGETRequest(server.URL)
	if err != nil {
		t.Errorf("estGateway.MakeRequest() error = %v", err)
		return
	}
	if got.StatusCode != wantStatusCode {
		t.Errorf("estGateway.MakeRequest() = %v, want %v", got, wantStatusCode)
	}
}
