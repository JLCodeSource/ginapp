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

	articleRoutes := router.Group("/article")
	{
		articleRoutes.Handle(http.MethodGet, "/view/:article_id", getArticle)
		articleRoutes.Handle(http.MethodGet, "/create", showArticleCreationPage)
		articleRoutes.Handle(http.MethodPost, "/create", createArticle)
	}

	
}