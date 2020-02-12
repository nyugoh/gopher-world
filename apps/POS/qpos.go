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
	r.Static("/assets", "./assets")

	// Load all templates
	r.LoadHTMLGlob("templates/*")

	// Register routes
	r.GET("/", qpos.IndexPage)
	r.GET("/stock/inventory", qpos.Fetchinventory)
	r.GET("/stock/stored-items", qpos.FetchStoredItems)
	r.GET("/stock/categories-suppliers", qpos.FetchCategoriesSuppliers)

	r.Run(":5000")
}
