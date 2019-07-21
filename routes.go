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

		userRoutes.Handle(http.MethodGet, "/register", register)
	}

	// Handle the article route
	router.Handle(http.MethodGet, "/article/view/:article_id", getArticle)
}