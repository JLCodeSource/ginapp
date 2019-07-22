package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
  }

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, ErrIDNotFound
}

func createNewArticle(title, content string) (*article, error) {
	//TODO investigate pointer to New Article

	lastID := len(articleList)

	a := article{ID: lastID+1, Title: title, Content: content}

	articleList = append(articleList, a)
	
	return &a, nil
}