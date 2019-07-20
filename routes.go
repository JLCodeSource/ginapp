package main

func initRoutes() {

	// Handle the index route
	router.Handle("GET", "/", showIndexPage)

	// Handle the article route
	router.Handle("GET", "/article/view/:article_id", getArticle)
}