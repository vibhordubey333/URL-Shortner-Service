package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	fs "github.com/vibhordubey333/URL-Shortner-Service/pkg/repository/filestorage"
	"github.com/vibhordubey333/URL-Shortner-Service/pkg/repository/redisdb"
	"github.com/vibhordubey333/URL-Shortner-Service/pkg/service"
	"log"
	"net/http"
	"strings"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	fmt.Println("Request received by CreateShortUrl")
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := service.KeyGenerator()
	//Storing in redis.

	redisdb.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	//Storing in file.
	fs.StoreInFile(shortUrl, creationRequest.LongUrl)

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
