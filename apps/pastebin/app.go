package main

import (
	. "github.com/nyugoh/pastebin/src/handler"
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippets/create/", CreateSnippet)
	mux.HandleFunc("/snippets/", ListSnippets)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)

}