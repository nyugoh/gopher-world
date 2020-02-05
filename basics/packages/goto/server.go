package UrlShortener

import (
	"fmt"
	"html/template"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	key:= r.URL.Path[1:]
	fmt.Println("Key: ", r.URL.Path[1:])
	if url := Store.Get(key); url != "" {
		fmt.Printf("Redirect: %s ==> %s", r.URL.Path[1:], url)
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.NotFound(w, r)
		return
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		// Send GET request
		fmt.Fprint(w, template.HTML(addForm))
		return
	}
	fmt.Println("Key", url)
	key := Store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}


var addForm = `
<html>
	<head>
		<title>Add Url</title>
	</head>
	<body>
		<form method="GET" action="/add">
			<label>Url</label>
			<input name="url" type="url"/>
			<input label="Add" type="submit"/>
		</form>
	</body>
</html>
`
