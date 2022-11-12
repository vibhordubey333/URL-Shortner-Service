package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateShortUrl(t *testing.T) {
	jsonText := "{\"long_url\":\"https://www.google.com\",\"user_id\":\"e0dba740-fc4b-4977-872c-d360239e6b10\"}"
	r := SetUpRouter()
	r.POST("/", CreateShortUrl)
	jsonRequest := strings.NewReader(jsonText)
	req, _ := http.NewRequest("POST", "/create-short-url", jsonRequest)

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			args: args{
				c: &gin.Context{
					Request: req,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateShortUrl(tt.args.c)
		})
	}
}

func TestHandleShortUrlRedirect(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleShortUrlRedirect(tt.args.c)
		})
	}
}
