package main

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"net/http"
	"strings"
)

func TestShowIndexPageUnauth(t *testing.T) {
	r := getRouter(true)

	r.Handle("GET", "/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	t.Run("returns the page title in the body", func (t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		page, err := ioutil.ReadAll(w.Body)
		
		title := "<title>Home Page</title>"

		assertStatus(t, w.Code, http.StatusOK)

		assertNoError(t, err)

		assertPageTitle(t, page, title)

	})
}

func TestGetArticle(t *testing.T) {
	r := getRouter(true)

	r.Handle("GET", "/article/view/:article_id", getArticle)

	t.Run("returns a single article", func(t *testing.T) {
		
		req, _ := http.NewRequest("GET", "/article/view/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusOK)
				
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error and got '%s'", err)
	}
}

func assertPageTitle(t *testing.T, page []byte, title string) {
	isTitle := strings.Index(string(page), title) > 0 
	if ! isTitle  {
		t.Errorf("title is not '%s' as expected", title)
	}

}