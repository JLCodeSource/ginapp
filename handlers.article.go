package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage (c *gin.Context) {

		articles := getAllArticles()

		// Call the HTML method of the Context to render a template	
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
				"payload": articles,
			},
		)
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
	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title": article.Title,
			"payload": article,
		},
	)	

}