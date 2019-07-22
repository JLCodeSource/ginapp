package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
)

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

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}