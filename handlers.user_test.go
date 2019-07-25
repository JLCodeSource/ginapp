package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.Handle(http.MethodGet, registerRoute, showRegistrationPage)

	req, _ := http.NewRequest(http.MethodGet, registerRoute, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	page, err := ioutil.ReadAll(w.Body)

	title := "<title>Register</title>"

	assertStatus(t, w.Code, http.StatusOK)
	assertNoError(t, err)
	assertPageContains(t, page, title)

}

func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	w := httptest.NewRecorder()

	r.Handle(http.MethodGet, loginRoute, showLoginPage)

	req, _ := http.NewRequest(http.MethodGet, loginRoute, nil)

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
	r.Handle(http.MethodPost, registerRoute, register)

	registrationPayload := getRegistrationPOSTPayload()

	req := getHeaders(t, http.MethodPost, registerRoute, registrationPayload)

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

	r.Handle(http.MethodPost, registerRoute, register)

	registrationPayload := getLoginPOSTPayload()

	req := getHeaders(t, http.MethodPost, registerRoute, registrationPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}

func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.Handle(http.MethodPost, loginRoute, performLogin)

	loginPayload := getRegistrationPOSTPayload()

	req := getHeaders(t, http.MethodPost, loginRoute, loginPayload)

	r.ServeHTTP(w, req)

	assertStatus(t, w.Code, http.StatusBadRequest)

	restoreLists()

}
