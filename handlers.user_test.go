package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"strings"
	"strconv"
)

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.Handle(http.MethodGet, "/u/register", showRegistrationPage)

	req, _ := http.NewRequest(http.MethodGet, "/u/register", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	page, err := ioutil.ReadAll(w.Body)

	title := "<title>Register</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, title)

}

func TestShowLoginPageUnauthenticated(t *testing.T){
	r := getRouter(true)
	w := httptest.NewRecorder()

	r.Handle(http.MethodGet, "/u/login", showLoginPage)

	req, _ := http.NewRequest(http.MethodGet, "u/login", nil)

	r.ServeHTTP(w, req)

	page, err := ioutil.ReadAll(w.Body)
	contains := "<title>Login</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, contains)

}

func TestRegisterUnauthenticated(t *testing.T) {
	saveLists()

	r := getRouter(true)
	w := httptest.NewRecorder()
	r.Handle(http.MethodPost, "/u/register", register)


	// TODO refactor out payload tests and headers
	registrationPayload := getRegistrationPOSTPayload()
	payload := strings.NewReader(registrationPayload)
	lenPayload := strconv.Itoa(len(registrationPayload))
	req, _ := http.NewRequest(http.MethodPost, "u/register", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", lenPayload)

	r.ServeHTTP(w, req)

	page, err := ioutil.ReadAll(w.Body)
	contains := "<title>Successful registration &amp; Login</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, contains)

	restoreLists()

}



func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	saveLists()
	r := getRouter(true)
	w := httptest.NewRecorder()

	r.Handle(http.MethodPost, "/u/register", register)

	// TODO refactor out payload tests and headers
	registrationPayload := getLoginPOSTPayload()
	payload := strings.NewReader(registrationPayload)
	lenPayload := strconv.Itoa(len(registrationPayload))
	req, _ := http.NewRequest(http.MethodPost, "/u/register", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", lenPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}


func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)
	
	r.Handle(http.MethodPost, "/u/login", performLogin)

	// TODO refactor out payload tests and headers
	loginPayload := getRegistrationPOSTPayload()
	payload := strings.NewReader(loginPayload)
	lenPayload := strconv.Itoa(len(loginPayload))
	req, _ := http.NewRequest(http.MethodPost, "/u/register", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", lenPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}

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