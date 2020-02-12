package main

import (
	"github.com/gin-gonic/gin"
	qpos "github.com/nyugoh/gopher-world/apps/POS/routes"
)

var (
	r *gin.Engine
)

func main() {
	r = gin.Default()

	// Server static files
	r.Static("/assets/", "./assets")

	// Load all templates
	r.LoadHTMLGlob("templates/*")

	// Register routes
	r.GET("/", qpos.IndexPage)

	r.Run(":5000")
}
