package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestUsernameAvailability(t *testing.T) {
	saveLists()
	newusername := "newuser"
	existingusername := "user1"
	
	assertUsernameAvailable(t, newusername, true)
	assertUsernameAvailable(t, existingusername, false)

	registerNewUser("newuser", "newpass")

	assertUsernameAvailable(t, newusername, false)

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	user := "newuser"

	u, err := registerNewUser(user, "newpass")

	assertNoError(t, err)
	assertUser(t, u, user)
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	t.Run("cannot register existing user" , func(t *testing.T){

		u, err := registerNewUser("user1", "pass1")

		assertError(t, err, ErrUsernameUnavailable)
		assertNilUser(t, u)
	})
	
	t.Run("cannot register user without valid password", func(t *testing.T){

		u, err := registerNewUser("newuser", "")

		assertError(t, err, ErrPasswordNotEmpty)
		assertNilUser(t, u)
	
	})

	restoreLists()
}

func TestUserValidity(t *testing.T) {

	user1 := "user1"
	user1Cap := "User1"
	user2 := "user2"
	pass1 := "pass1"
	pass1Cap := "Pass1"
	empty := ""

	validTests := []struct {
		name string
		pass string
		valid bool
	}{
		{name: user1, pass: pass1, valid: true},
		{name: user2, pass: pass1, valid: false},
		{name: user1, pass: empty, valid: false},
		{name: empty, pass: pass1, valid: false},
		{name: user1Cap, pass: pass1, valid: false},
		{name: user1, pass: pass1Cap, valid: false},
	}

	for _, tt := range validTests {
		assertUserValid(t, tt.name, tt.pass, tt.valid)
	}
}

func TestLoginUnauthenticated(t *testing.T){
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.Handle(http.MethodPost, loginRoute, performLogin)

	loginPayload := getLoginPOSTPayload()

	req := getHeaders(t, http.MethodPost, loginRoute, loginPayload)

	r.ServeHTTP(w, req)

	page, err := ioutil.ReadAll(w.Body)
	contains := "<title>Successful Login</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, contains)

	restoreLists()

}
