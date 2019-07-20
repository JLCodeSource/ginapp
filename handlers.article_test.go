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

		got := w.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d, want status %d", got, want)
		}

		page, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Errorf("expected no error and got '%s'", err)
		}
		title := "<title>Home Page</title>"
		pageTitle := strings.Index(string(page), title) > 0 
		if pageTitle != true {
			t.Errorf("title is not '%s' as expected", title)
		}

	})
}

func TestGetArticle(t *testing.T) {
	r := getRouter(true)

	r.Handle("GET", "/article/view/:article_id", getArticle)

	t.Run("returns a single article", func(t *testing.T) {
		
		req, _ := http.NewRequest("GET", "/article/view/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		want := 200
		got := w.Code
		if got != want {
			t.Errorf("got status %d, want status %d", got, want)
			t.Errorf("%s", w.Body)
		}	
		
	})
}