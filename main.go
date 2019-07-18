package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	
	// Set router as Gin default 
	router = gin.Default()

	// Process templates at start to avoid re-reading
	router.LoadHTMLGlob("templates/*")
	
	// Define route for index page
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template	
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	// Serve application
	router.Run()
}