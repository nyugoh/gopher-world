package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Paste Bin"))
}

func ListSnippets(w http.ResponseWriter, r* http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Please provide an id", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "List of snippet for : %d", id)
}

func CreateSnippet(w http.ResponseWriter, r* http.Request)  {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}
	w.Write([]byte("Create snippets"))
}