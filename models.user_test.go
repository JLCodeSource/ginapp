package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"strconv"
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

	//TODO tabularize tests & refactor with assertValid

	if !isUserValid(user1, pass1) {
		t.Errorf("expected user '%s' and pass '%s' to validate but they did not", 
			user1, pass1)
	}

	if isUserValid(user2, pass1) {
		t.Errorf("expected user '%s' and pass '%s' to be invalid and they weren't", 
			user2, pass1)
	}

	if isUserValid(user1, empty) {
		t.Errorf("expected user '%s' and pass '%s' to be invalid and they weren't", 
			user1, empty)
	}

	if isUserValid(empty, pass1) {
		t.Errorf("expected user '%s' and pass '%s' to be invalid and they weren't",
			empty, pass1)
	}

	if isUserValid(user1Cap, pass1) {
		t.Errorf("expected user '%s' and pass '%s' to be invalid and they weren't",
			user1Cap, pass1)
	}

	if isUserValid(user1, pass1Cap) {
		t.Errorf("expected user '%s' and pass '%s' to be invalid and they weren't",
			user1, pass1Cap)
	}

}

func TestLoginUnauthenticated(t *testing.T){
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.Handle(http.MethodPost, "/u/login", performLogin)

	// TODO refactor out payload tests and headers

	loginPayload := getLoginPOSTPayload()
	payload := strings.NewReader(loginPayload)
	loginlen := strconv.Itoa(len(loginPayload))
	req, _ := http.NewRequest(http.MethodPost, "/u/login", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", loginlen)

	r.ServeHTTP(w, req)

	page, err := ioutil.ReadAll(w.Body)
	contains := "<title>Successful Login</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, contains)

	restoreLists()

}
