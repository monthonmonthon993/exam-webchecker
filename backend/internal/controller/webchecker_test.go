package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"exam.com/webchecker/internal/core"
	"exam.com/webchecker/internal/gateway"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_webcheckerController_CheckAvalibleWebsites(t *testing.T) {
	type fields struct {
		webCheckerCore core.WebCheckerCore
	}
	tests := []struct {
		name    string
		fields  fields
		content string
		want    string
	}{
		{
			name:    "test_CheckAvalibleWebsites_success",
			content: "www.google.com,\nwww.facebook.com",
			want:    `{"website_status_list":[{"name":"www.facebook.com","ok":true,"err_msg":null},{"name":"www.google.com","ok":true,"err_msg":null}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &webcheckerController{
				webCheckerCore: core.NewWebCheckerCore(gateway.NewMockRESTGateway(nil)),
			}
			r := setupRouter(c.CheckAvalibleWebsites)

			w := httptest.NewRecorder()

			req := newfileUploadRequest(t, tt.content, "/webchecker/websites", "csv", "test.csv")
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			wantStruct := response{}
			gotStruct := response{}

			assert.NoError(t, json.Unmarshal([]byte(tt.want), &wantStruct))
			assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &gotStruct))

			assert.Equal(t, len(wantStruct.WebsiteStatusList), len(gotStruct.WebsiteStatusList))
			for _, wantBody := range wantStruct.WebsiteStatusList {
				found := false
				for _, gotBody := range wantStruct.WebsiteStatusList {
					if wantBody == gotBody {
						found = true
					}
				}
				assert.Equal(t, true, found)
			}
		})
	}
}

// Creates a fake router for this testing
func setupRouter(fn func(c *gin.Context)) *gin.Engine {
	r := gin.Default()
	r.POST("/webchecker/websites", fn)
	return r
}

// Creates a new file upload to fake http request
func newfileUploadRequest(t *testing.T, content string, uri, paramName, path string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile(paramName, uri)
	assert.NoError(t, err)

	f := strings.NewReader(content)
	_, err = io.Copy(fw, f)
	assert.NoError(t, err)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, uri, bytes.NewReader(body.Bytes()))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}
