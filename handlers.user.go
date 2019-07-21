package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func showRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	_, err := registerNewUser(username, password)

	// If username and password combination is invalid,
	// show the error message on the login page
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle": "Registration Failed",
			"ErrorMessage": err.Error()})
	} else {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")
	}
	
}

func generateSessionToken() string {
	// We're using a random 16 char string as session token
	// Not a secure way to gen session tokens - Not for Prod
	return strconv.FormatInt(rand.Int63(), 16)
}