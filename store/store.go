package store

import "fmt"

// Item properties
type Item struct {
	Name string
	Price float64
	Code string
}

// Specific item
type storeItem map[string]Item

type Store struct {
	Items []storeItem
	Total int
}

var (
	Stock Store
)

func init() {
	addToStore(&Item{"A4", 100, "M0001"})
	addToStore(&Item{"Pencil", 20, "M0002"})
}

func AddStock() {
	var (
		name string
		price float64
		code string
	)
	fmt.Println("*****************")
	fmt.Println("Adding item")
	fmt.Printf("Enter item name ==>")
	fmt.Scanf("%s", &name)
	fmt.Printf("\nEnter item code ==>")
	fmt.Scanf("%s", &code)
	fmt.Printf("\nEnter item price ==>")
	fmt.Scanf("%f", &price)
	addToStore(&Item{name, price, code})
}

func addToStore(newItem *Item)  {
	item := make(storeItem)
	item[newItem.Code] = *newItem
	stockItems := make([]storeItem, 0)
	stockItems = append(stockItems, item)
	if Stock.Items == nil {
		Stock = Store{stockItems, 1}
	} else {
		Stock.Items = append(Stock.Items, item)
		Stock.Total += Stock.Total
	}
}
