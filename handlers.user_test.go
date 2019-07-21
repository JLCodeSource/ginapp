package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
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

	registrationPayload := getRegistrationPOSTPayload()
	
	req := getHeaders(t, http.MethodPost, "u/register", registrationPayload)

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

	registrationPayload := getLoginPOSTPayload()

	req := getHeaders(t, http.MethodPost, "/u/register", registrationPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}


func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)
	
	r.Handle(http.MethodPost, "/u/login", performLogin)

	loginPayload := getRegistrationPOSTPayload()

	req := getHeaders(t, http.MethodPost, "/u/login", loginPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}

