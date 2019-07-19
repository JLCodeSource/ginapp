package main

func initRoutes() {

	// Handle the index route
	router.Handle("GET", "/", showIndexPage)
}