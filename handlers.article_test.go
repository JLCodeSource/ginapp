package main

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"encoding/xml"
	//"strings"
	//"strconv"
)

func TestShowIndexPageUnauth(t *testing.T) {
	r := getRouter(true)

	r.Handle(http.MethodGet, "/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

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
/* 	t.Run("returns a not found on a non-integer id", func(t *testing.T) {		
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

	}) */
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

func TestArticleCreationAuthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.Handle(http.MethodPost, "/article/create", createArticle)

	articlePayload := getArticlePOSTPayload()

	req := getHeaders(t, http.MethodPost, createRoute, articlePayload)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	
	r.ServeHTTP(w, req)

	page, err := ioutil.ReadAll(w.Body)
	contains := "<title>Submission Successful</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	//TODO add assertAuthenticated
	assertPageContains(t, page, contains)

	restoreLists()

}
