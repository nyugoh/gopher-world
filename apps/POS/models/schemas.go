package qpos

import (
	"time"
)

type StoredItems struct {
	ID           int64  `json:"id"`
	Categoryid   int64  `json:"categoryid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ReorderPoint int64  `json:"reorderpoint"`
	StockType    int64  `json:"stockType"`
	Taxid        int64  `json:"taxid"`
}

type StoredCategories struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type StoredSuppliers struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TaxRules struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Inclusive bool          `json:"inclusive"`
	Mode      string        `json:"mode"`
	Base      []interface{} `json:"base"`
	Locations []interface{} `json:"locations"`
}

type TaxItems struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"name"`
	Value      string `json:"name"`
	Multiplier string `json:"name"`
}

type StockItems struct {
	ID           int       `json:"id"`
	StoredItemId int       `json:"stockinventoryid"`
	SupplierId   int       `json:"supplierId"`
	StockLevel       int   `json:"stocklevel"`
	ExpiryDate   time.Time `json:"expiryDate"`
	Cost         float64   `json:"cost"`
	Price        float64   `json:"price"`
	WholePrice   float64   `json:"wprice"`
	Code         string    `json:"code"`
	Data         string    `json:"data"`
	LocationId   int       `json:"locationId"`
	InvoiceNo    string    `json:"inventoryNo"`
	StockType    int       `json:"stockType"`
}

type Inventory struct {
	ID           int       `json:"id"`
	StoredItemId int       `json:"storedItemId"`
	SupplierId   int       `json:"supplierId"`
	Amount       float64   `json:"amount"`
	ExpiryDate   time.Time `json:"expiryDate"`
	Cost         float64   `json:"cost"`
	Price        float64   `json:"price"`
	WholePrice   float64   `json:"wprice"`
	Code         string    `json:"code"`
	Data         string    `json:"data"`
	LocationId   int       `json:"locationId"`
	InvoiceNo    string    `json:"invoiceNo"`
	StockType    int       `json:"stockType"`
}

type InventoryHistory struct {
	ID           int64 `json:"id"`
	InventoryId  int   `json:"inventoryId"`
	LocationId   int   `json:"locationId"`
	ActionTypeId int   `json:"actionTypeId"`
	StockLevel   int   `json:"stockLevel"`
}
