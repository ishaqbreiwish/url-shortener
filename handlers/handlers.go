package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishaqbreiwish/go-url-shortener/shortener"
	"github.com/ishaqbreiwish/go-url-shortener/store"
)

// Defines what incoming post request should look like
// and the information that needs to be included in the JSONs
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	// Reads the JSON data sent in the request body and tries to map it to the creationRequest struct
	// If not possible (meaning JSON isnt in right format) then throws an error
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// uses the defined functions to shorten the links then save the url mappings
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl") // Extracts the dynamic URL parameter shortUrl from the request, defined in main.go
	initialUrl := store.RetrieveInitialUrl(shortUrl)

	// code 302 is used to indicate that the redirect is temp, and that the browser will not permanently
	// change the URL in it history
	c.Redirect(302, initialUrl) // Redirects the user to the original URL (initialUrl) with an HTTP 302 Found status code.

}
