package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	loginRoute    = "/u/login"
	registerRoute = "/u/register"
	createRoute   = "/article/create"
)

var dummyArticleList []article

var dummyUserList []user

// This function is used for setup before executing the test functions
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
		r.Use(setUserStatus())
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

// TODO refactor get...Payload

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func getArticlePOSTPayload() string {
	params := url.Values{}
	params.Add("ID", "3")
	params.Add("Title", "Article 3")
	params.Add("Content", "Article 3 body")
	return params.Encode()
}

func getHeaders(t *testing.T, method, route, payload string) *http.Request {

	sPayload := strings.NewReader(payload)
	lenPayload := strconv.Itoa(len(payload))

	req, _ := http.NewRequest(method, route, sPayload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", lenPayload)

	return req
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

func assertUser(t *testing.T, got *user, want string) {
	t.Helper()
	if got.Username != want {
		t.Errorf("wanted username '%s' but got '%s'", got.Username, want)
	}
}

func assertNilUser(t *testing.T, user *user) {
	t.Helper()
	if user != nil {
		t.Errorf("expected nil response, but got username '%s'", user.Username)
	}
}

func assertUsernameAvailable(t *testing.T, username string, available bool) {
	t.Helper()
	got := isUsernameAvailable(username)
	want := available
	if got != want {
		t.Errorf("wanted username '%s' available to be %t, but got %t", username, want, got)
	}
}

func assertUserValid(t *testing.T, username, password string, valid bool) {
	t.Helper()
	got := isUserValid(username, password)
	want := valid
	if got != want {
		t.Errorf("wanted user '%s' with pass '%s' validity to be %t but they were %t",
			username, password, want, got)
	}
}

func assertPageContains(t *testing.T, page []byte, content string) {
	t.Helper()
	isContent := strings.Index(string(page), content) > 0
	if !isContent {
		t.Errorf("page does not contain '%s' as expected", content)
	}

}

func assertNumberOfArticles(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d articles, wanted %d", got, want)
	}
}

func assertArticleTitle(t *testing.T, a *article, title string) {
	t.Helper()
	got := a.Title
	want := title
	if got != want {
		t.Errorf("got '%s' title, wanted '%s'", got, want)
	}
}

func assertArticleContent(t *testing.T, a *article, content string) {
	t.Helper()
	got := a.Content
	want := content
	if got != want {
		t.Errorf("got '%s' content, wanted '%s'", got, want)
	}
}
