package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO - turn tests below into a table test

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setLoggedIn(false), ensureLoggedIn(), func(c *gin.Context) {
		t.Errorf("not logged in so expected not be seen as logged in")
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusUnauthorized)
}
func TestEnsureLoggedInAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setLoggedIn(true), ensureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusOK)
}

func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setLoggedIn(true), ensureNotLoggedIn(), func(c *gin.Context) {
		t.Errorf("logged in, but seen as not logged in")
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusUnauthorized)

}

func TestEnsureNotLoggedInUnauthenticted(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setLoggedIn(false), ensureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusOK)
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Errorf("expected to be set as authenticated")
		}
	})

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusOK)
}

func TestSetUserStatusUnauthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	r.Handle(http.MethodGet, "/", setUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if exists && loggedInInterface.(bool) {
			t.Errorf("expected to be set as unauthenticated")
		}
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusOK)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
