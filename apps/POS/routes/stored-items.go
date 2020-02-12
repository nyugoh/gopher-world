package qpos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	qpos2 "github.com/nyugoh/gopher-world/apps/POS/models"
	qpos "github.com/nyugoh/gopher-world/apps/POS/utils"
	"net/http"
)

func FetchStoredItems(c *gin.Context) {
	db := qpos.DbConnect()
	defer db.Close()
	var items []qpos2.StoredItems
	db.Find(&items)
	for _, item:= range items {
		fmt.Println("Cat Id",item.CategoryId)
		fmt.Println("Reorder", item)
	}
	c.HTML(http.StatusOK, "stored-items.html", gin.H{
		"title": "Stored Items",
		"payload": items,
	})
}

func FetchCategoriesSuppliers(c *gin.Context) {
	db := qpos.DbConnect()
	defer db.Close()
	var categories []qpos2.StoredCategories
	var suppliers []qpos2.StoredSuppliers
	db.Find(&categories)
	db.Find(&suppliers)
	fmt.Println(categories)
	fmt.Println(suppliers)
	c.HTML(http.StatusOK, "categories-suppliers.html", gin.H{
		"title": "Categories & Suppliers",
		"categories": categories,
		"suppliers": suppliers,
	})
}

func Fetchinventory(c *gin.Context)  {
	db :=qpos.DbConnect()
	defer db.Close()
	var inventory []qpos2.StockItems
	db.Find(&inventory)

	c.HTML(http.StatusOK, "inventory.html", gin.H{
		"title": "Stock",
		"inventory": inventory,
	})
}