package main

import (
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"strings"
)


var dummyArticleList []article

var dummyUserList []user

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
	dummyUserList = userList
	dummyArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	userList = dummyUserList
	articleList = dummyArticleList
}


func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error and got '%s'", err)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected error but didn't get one")
	}
	if got != want {
		t.Errorf("got error '%s', expected '%s'", got, want)
	}
}

func assertPageContains(t *testing.T, page []byte, content string) {
	t.Helper()
	isContent := strings.Index(string(page), content) > 0 
	if ! isContent  {
		t.Errorf("page does not contain '%s' as expected", content)
	}

}