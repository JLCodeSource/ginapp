package main

import (
	"github.com/gin-gonic/gin"
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