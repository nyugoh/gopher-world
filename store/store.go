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
	option := 1
	fmt.Println("\nOptions 0: Back 1: Add other")
	fmt.Scanf("%d", &option)
	 for {
		 if option == 0 {
			 break;
		 }
		 if option == 1 {
			 continue
		 }
		 var (
			 name string
			 price float64
			 code string
		 )
		 fmt.Println("\n*****************")
		 fmt.Println("   Adding item")
		 fmt.Println("*****************")
		 fmt.Printf("Enter item name ==>")
		 fmt.Scanf("%s", &name)
		 fmt.Printf("\nEnter item code ==>")
		 fmt.Scanf("%s", &code)
		 fmt.Printf("\nEnter item price ==>")
		 fmt.Scanf("%f", &price)

		 addToStore(&Item{name, price, code})

		 fmt.Println("\nOptions 0: Back 1: Add other")
		 fmt.Scanf("%d", &option)
	 }
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

func ListStock()  {
	fmt.Println("\n\n\t******** Store items *********")
	fmt.Println("\n\t--------------------------------")
	fmt.Println("\t| # |   Name   | Code | Price  |")
	fmt.Println("\t--------------------------------")
	for idx, item := range Stock.Items {
		for code, sItem := range item {
			fmt.Printf("\t|%3d|%10s|%6s|%8.2f|\n", idx, sItem.Name, code, sItem.Price)
		}
		fmt.Println("\t--------------------------------")
	}
	fmt.Println("\n\t---------------*----------------")
}