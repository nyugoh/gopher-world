package routes

import (
	"github.com/nyugoh/gopher-world/templates/rest-api/utils"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status": "200",
		"message": "Welcome to andromeda",
	})
}

