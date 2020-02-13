package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nyugoh/gopher-world/templates/rest-api/routes"
	"github.com/nyugoh/gopher-world/templates/rest-api/utils"
	"net/http"
	"os"

)

func main() {
	utils.InitLogger()
	utils.Log("Starting app...")

	if err:= godotenv.Load(); err != nil {

	}

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeRoute)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	port := ":5000"
	if len(os.Getenv("PORT")) != 0 {
		port = ":"+os.Getenv("PORT")
	}

	if err := http.ListenAndServe(port, loggedRouter); err != nil {
		utils.Log("Error: Unable to start app: ", err.Error())
	} else {
		utils.Log("Started app successfully...")
	}
}
