package main

import (
	"os"
	"testing"
	"github.com/gin-gonic/gin"
)

var dummyArticleList []article

// This funciton is used for setup before executing the test functions
func TestMain(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// This function is used to store the main lists into the temporary
// list for testing
func saveLists() {
	dummyArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	articleList = dummyArticleList
}