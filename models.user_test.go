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
	
	//TODO refactor with asserts

	if !isUsernameAvailable(newusername) {
		t.Errorf("expected username '%s' to be available, but it is not", newusername)
	}

	if isUsernameAvailable(existingusername) {
		t.Errorf("expected username '%s' to be unavailable, but it is not", existingusername)
	}

	registerNewUser("newuser", "newpass")
	if isUsernameAvailable(newusername) {
		t.Errorf("expected username '%s' to be unavailable, but it is not", newusername)
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	//TODO refactor with asserts

	u, err := registerNewUser("newuser", "newpass")
	empty := ""

	if err != nil {
		t.Errorf("did not expect an error but got '%s'", err)
	}

	if u.Username == empty {
		t.Errorf("wanted username '%s' but got '%s'", u.Username, empty)
	}
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	//TODO refactor with asserts

	u, err := registerNewUser("user1", "pass1")

	if err == nil {
		t.Errorf("expected error but got none")
	}

	if u != nil {
		t.Errorf("expected nil response, but got '%s'", u)
	}

	u, err = registerNewUser("newuser", "")

	
	if err == nil {
		t.Errorf("expected error but got none")
	}

	if u != nil {
		t.Errorf("expected nil response, but got '%s'", u)
	}

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
