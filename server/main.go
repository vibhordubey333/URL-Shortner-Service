package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vibhordubey333/URL-Shortner-Service/pkg/handler"
	"github.com/vibhordubey333/URL-Shortner-Service/pkg/repository/redisdb"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API\n1. /create-short-url\n 2. /:shortUrl",
		})
	})

	r.POST("/api/shorturls", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/api/shorturls/:shortUrl", func(c *gin.Context) {
		handler.FetchOriginalURL(c)
	})
	// Note store initialization happens here
	redisdb.InitializeStore()
	errorResponse := r.Run(":7777")
	if errorResponse != nil {
		panic(fmt.Sprintf("Error while starting URL-Shortner-Server%v", errorResponse))
	}

}
