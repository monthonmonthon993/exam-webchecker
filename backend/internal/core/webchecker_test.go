package core

import (
	"testing"

	"exam.com/webchecker/internal/entities"
	"exam.com/webchecker/internal/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_webCheckerCore_CheckAvalibleWebsite(t *testing.T) {
	type fields struct {
		websiteGateway gateway.RESTGateway
	}
	type args struct {
		website string
		ch      chan entities.WebsiteStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    entities.WebsiteStatus
	}{
		{
			name: "test_CheckAvalibleWebsite_success",
			args: args{
				website: "https://facebook.com",
			},
			want: entities.WebsiteStatus{
				Name: "https://facebook.com",
				OK:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &webCheckerCore{
				websiteGateway: gateway.NewMockRESTGateway(nil),
			}
			tt.args.ch = make(chan entities.WebsiteStatus)
			got, err := w.CheckAvalibleWebsite(tt.args.website)
			if (err != nil) != tt.wantErr {
				t.Errorf("webCheckerCore.CheckAvalibleWebsite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)

		})
	}
}
