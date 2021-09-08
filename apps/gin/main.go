package main

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func init() {
	articles = []Article{
		{
			Id:      1,
			Title:   "Article 1",
			Content: "Content 1",
		},
		{
			Id:      2,
			Title:   "Article 2",
			Content: "Content 2",
		},
		{
			Id:      3,
			Title:   "Article 3",
			Content: "Content 3",
		},
	}
}

func main() {
	r = gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", IndexPage)
	r.GET("/article/:id", ShowArticle)

	r.Run(":5000")
}
