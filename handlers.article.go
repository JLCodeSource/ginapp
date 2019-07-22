package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage (c *gin.Context) {

		articles := getAllArticles()

		// Call the HTML method of the Context to render a template	
		render(c, gin.H{
				"title": "Home Page",
				"payload": articles,
			}, "index.html")
}

func getArticle (c *gin.Context) {

	// Check for valid article ID in URL
	articleID, err := strconv.Atoi(c.Param("article_id"))
	// if not valid article ID in URL report Not Found
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		
	}

	// Check whether article exists
	article, err := getArticleByID(articleID)
	// if article not found report article not found error
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// display article
	render(c, gin.H{
			"title": article.Title,
			"payload": article,
		}, "article.html")	

}

func createArticle (c *gin.Context) { }