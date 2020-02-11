package main

type Article struct {
	Id      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	articles []Article
)

func GetAllArticles() (arts []Article) {
	arts = articles
	return
}
