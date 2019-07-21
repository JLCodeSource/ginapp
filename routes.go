package main

import (
	"net/http"
)

func initRoutes() {

	// Handle the index route
	router.Handle(http.MethodGet, "/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.Handle(http.MethodGet, "/register", showRegistrationPage)
		userRoutes.Handle(http.MethodPost, "/register", register)
		userRoutes.Handle(http.MethodGet, "/login", showLoginPage)
		userRoutes.Handle(http.MethodPost, "/login", performLogin)
		userRoutes.Handle(http.MethodGet, "/logout", logout)

	}

	// Handle the article route
	router.Handle(http.MethodGet, "/article/view/:article_id", getArticle)
}