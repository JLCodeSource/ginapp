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

	// Initialize the routes
	initRoutes()

	// Serve application
	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {

	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
