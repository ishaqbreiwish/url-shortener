package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/ishaqbreiwish/go-url-shortener/handlers"
	"github.com/ishaqbreiwish/go-url-shortener/store"
)

func main() {
	r := gin.Default() // creates a new gin engine and saves it to r

	r.GET("/", func(c *gin.Context) { // checks the root URL
		// gin.Context provides everything we need to handle
		// incoming HTTP requests and send a response back to the client
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API", // creates a map
		})
	}) // Closing r.GET properly here

	// listens for POST requests made to create short url
	// when a user sends a POST request to the URL CreateShortUrl is called
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Listens to Get requests sent to URLs matching the pattern /:shortUrl
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808") // starts the gin web server at port 9808
	if err != nil {       // if r.Run returned an error, the server failed to start
		// if there is an error, we raise a panic which terminates the program
		// fmt.Sprintf creates a formatted error message
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
