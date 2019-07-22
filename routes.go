package main

import (
	"net/http"
)

func initRoutes() {

	router.Use(setUserStatus())
	// Handle the index route
	router.Handle(http.MethodGet, "/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.Handle(http.MethodGet, "/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.Handle(http.MethodPost, "/register", ensureNotLoggedIn(), register)
		userRoutes.Handle(http.MethodGet, "/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.Handle(http.MethodPost, "/login", ensureNotLoggedIn(), performLogin)
		userRoutes.Handle(http.MethodGet, "/logout", ensureLoggedIn(), logout)

	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.Handle(http.MethodGet, "/view/:article_id", getArticle)
		articleRoutes.Handle(http.MethodGet, "/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.Handle(http.MethodPost, "/create", ensureLoggedIn(), createArticle)
	}

	
}