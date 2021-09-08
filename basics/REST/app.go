package main

import (
	"github.com/gorilla/mux"
	"github.com/nyugoh/gopher-world/basics/REST/rest"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	rest.CreateTable()
	r.HandleFunc("/", rest.Index).Methods("GET")
	r.HandleFunc("/welcom/{name}", rest.Greet).Methods("GET")

	postsRouter := r.Path("/posts").Subrouter()
	postsRouter.HandleFunc("", rest.AllPosts).Methods("GET")
	postsRouter.HandleFunc("", rest.AddPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
