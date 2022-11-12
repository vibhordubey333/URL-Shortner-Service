package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"vibhordubey333/URL-Shortner-Service/pkg/repository/redisdb"
	"vibhordubey333/URL-Shortner-Service/pkg/service"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := service.KeyGenerator()
	redisdb.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	shortUrl = strings.Replace(shortUrl, ":", "", -1)
	log.Println("Request Received By HandleShortUrlRedirect: ", shortUrl)

	initialUrl := redisdb.RetrieveInitialUrl(shortUrl)
	c.JSON(200, gin.H{
		"LongURL": initialUrl,
	})
}
