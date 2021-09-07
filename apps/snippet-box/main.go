package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}


func showSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Showing a specific snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Creating a snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	if err := http.ListenAndServe(":http", mux); err != nil {
		log.Fatal("Error starting server::", err.Error())
	}
}