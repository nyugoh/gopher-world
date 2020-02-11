package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexPage(c *gin.Context) {
	allArticles := GetAllArticles()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home page",
		"payload": allArticles,
	})
}

func ShowArticle(c *gin.Context) {
	if articleId, err := strconv.Atoi(c.Param("id")); err == nil {
		if article, err := getArticle(articleId); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func getArticle(articleId int) ( *Article, error) {

	for _, article := range GetAllArticles() {
		if article.Id == articleId {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}
