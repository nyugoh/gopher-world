package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	fmt.Fprintf(w, "Hello %s from Golang", vars["name"])
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world from Go")
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start of the middleware.")
		handler.ServeHTTP(w, r)
		fmt.Println("End of the middleware.")
	})

}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := CreateDBConnection()
	limit := r.URL.Query().Get("limit")
	fmt.Println(limit)
	query := "SELECT id, title, description, meta, content, image_cover, created_at FROM posts " + limit
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Err occured when adding post", err)
	}
	fmt.Println(rows.Columns())
	var posts []Post
	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.Title, &post.Description, &post.Meta, &post.Content, &post.ImageCover, &post.CreateAt)
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Print(posts)
	fmt.Fprint(w, posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	post := Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO posts (title, description, meta, content, image_cover, created_at) VALUES(?,?, ?,?,?, ?)"
	db := CreateDBConnection()
	defer db.Close()
	result, err := db.Exec(query, post.Title, post.Description, post.Meta, post.Content, post.ImageCover, time.Now())
	if err != nil {
		fmt.Fprintf(w, "Error when adding: %v", err)
		return
	}
	id, _ := result.LastInsertId()

	fmt.Print(post)
	fmt.Fprintf(w, "Inserted %v with id #%d", post, id)
}
