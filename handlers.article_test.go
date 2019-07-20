package main

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
	"encoding/xml"
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
		assertPageContains(t, page, title)

	})
}

func TestGetArticle(t *testing.T) {
	r := getRouter(true)

	r.Handle(http.MethodGet, "/article/view/:article_id", getArticle)

	t.Run("returns a single article", func(t *testing.T) {
		
		req, _ := http.NewRequest(http.MethodGet, "/article/view/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		page, err := ioutil.ReadAll(w.Body)

		title := "<title>Article 1</title>"
		body := "<p>Article 1 body</p>"

		assertStatus(t, w.Code, http.StatusOK)
		assertNoError(t, err)
		assertPageContains(t, page, title)
		assertPageContains(t, page, body)

	})
	t.Run("returns a not found on a non-integer id", func(t *testing.T) {		
		req, _ := http.NewRequest(http.MethodGet, "/article/view/asdasda", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
			
		assertStatus(t, w.Code, http.StatusNotFound)
	})
 	t.Run("returns a not found and error on non-existent article", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/article/view/3", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		_, err := getArticleByID(3)
		want := ErrIDNotFound
		assertStatus(t, w.Code, http.StatusNotFound)
		assertError(t, err, want)

	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	r.Handle(http.MethodGet, "/", showIndexPage)

	t.Run("returns multiple articles", func(t *testing.T) {
		
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		page, err := ioutil.ReadAll(w.Body)

		var articles []article
		err = json.Unmarshal(page, &articles)

		assertStatus(t, w.Code, http.StatusOK)
		assertNoError(t, err)
		if ! (len(articles) >= 2) {
			t.Errorf("expected 2 or more articles got '%d'", len(articles))
		}

	})
}

func TestArticleXML(t *testing.T){
	r := getRouter(true)

	r.Handle(http.MethodGet, "/article/view/:article_id", getArticle)

	t.Run("returns multiple articles", func(t *testing.T) {
		
		req, _ := http.NewRequest(http.MethodGet, "/article/view/1", nil)
		req.Header.Add("Accept", "application/xml")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		page, err := ioutil.ReadAll(w.Body)

		var a article
		err = xml.Unmarshal(page, &a)

		assertStatus(t, w.Code, http.StatusOK)
		assertNoError(t, err)
		if ! (a.ID == 1 && len(a.Title) >= 0) {
			t.Errorf("expected id 1 and title len >= 0 but got '%d' & '%d'", a.ID, len(a.Title))
		}

	})
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