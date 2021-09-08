package main

import (
	"github.com/nyugoh/tuts/goto"
	"net/http"
)

func main() {
	UrlShortener.Store = UrlShortener.GenerateNewUrlStore()

	http.HandleFunc("/", UrlShortener.Redirect)
	http.HandleFunc("/add", UrlShortener.Add)

	http.ListenAndServe(":8080", nil)
}
